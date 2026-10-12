package com.max1874.ferry
import android.content.ContentResolver
import android.net.Uri
import android.provider.DocumentsContract
import android.provider.OpenableColumns
import java.io.ByteArrayOutputStream
import java.io.InputStream
import java.io.OutputStream

/**
 * A picked file is read from [uri] when it is sent. A pasted image carries its
 * [bytes] instead, because its read grant does not last that long.
 */
data class SelectedContent(
    val uri: String,
    val name: String,
    val mediaType: String,
    val size: Long?,
    val bytes: ByteArray? = null,
)

interface ContentStore {
    suspend fun describe(uri: String): SelectedContent
    fun open(uri: String): InputStream
    fun create(uri: String): OutputStream
    fun delete(uri: String): Boolean

    /**
     * Reads a selection into memory now. The system revokes an app's read
     * grant on a clipboard item as soon as anything else is copied, so a pasted
     * image cannot wait until it is sent.
     */
    suspend fun snapshot(uri: String): SelectedContent {
        val described = describe(uri)
        // Near the 64 MB limit the buffer and its copy can exhaust the heap;
        // that is an Error, which the caller's Exception handler would miss.
        val bytes = try { open(uri).use { input ->
            val buffer = ByteArrayOutputStream()
            val chunk = ByteArray(256 * 1024)
            while (true) {
                val count = input.read(chunk)
                if (count < 0) break
                if (buffer.size() + count > FerryJson.MAX_FILE_BYTES) {
                    throw FerryApiException(413, "payload_too_large", "Files must be 64 MB or smaller.")
                }
                buffer.write(chunk, 0, count)
            }
            buffer.toByteArray()
        } } catch (_: OutOfMemoryError) {
            throw FerryApiException(413, "payload_too_large", "This image is too large to paste. Attach it from Files instead.")
        }
        return described.copy(size = bytes.size.toLong(), bytes = bytes)
    }
}

class AndroidContentStore(private val resolver: ContentResolver) : ContentStore {
    override suspend fun describe(uri: String): SelectedContent {
        val value = Uri.parse(uri)
        var name: String? = null
        var size: Long? = null
        resolver.query(value, arrayOf(OpenableColumns.DISPLAY_NAME, OpenableColumns.SIZE), null, null, null)?.use { cursor ->
            if (cursor.moveToFirst()) {
                val nameIndex = cursor.getColumnIndex(OpenableColumns.DISPLAY_NAME)
                val sizeIndex = cursor.getColumnIndex(OpenableColumns.SIZE)
                if (nameIndex >= 0 && !cursor.isNull(nameIndex)) name = cursor.getString(nameIndex)
                if (sizeIndex >= 0 && !cursor.isNull(sizeIndex)) size = cursor.getLong(sizeIndex)
            }
        }
        val displayName = name?.takeIf { it.isNotBlank() } ?: value.lastPathSegment ?: "file"
        val mediaType = resolver.getType(value)?.takeIf(::safeMediaType) ?: "application/octet-stream"
        return SelectedContent(uri, displayName, mediaType, size)
    }

    override fun open(uri: String): InputStream = resolver.openInputStream(Uri.parse(uri))
        ?: throw IllegalArgumentException("The selected file cannot be opened.")

    override fun create(uri: String): OutputStream = resolver.openOutputStream(Uri.parse(uri), "wt")
        ?: throw IllegalArgumentException("The destination file cannot be opened.")

    override fun delete(uri: String): Boolean = runCatching {
        DocumentsContract.deleteDocument(resolver, Uri.parse(uri))
    }.getOrDefault(false)

    private fun safeMediaType(value: String): Boolean = value.length <= 127 &&
        value.count { it == '/' } == 1 && value.none { it <= ' ' || it == '\u007f' }
}
