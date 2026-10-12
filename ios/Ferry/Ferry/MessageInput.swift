import SwiftUI
import UIKit
import UniformTypeIdentifiers

/// An image taken off the pasteboard, in the encoding Ferry will upload.
struct PastedImage: Equatable {
    let data: Data
    let mediaType: String
    let filenameExtension: String

    /// Formats every Ferry client can preview are kept byte for byte, so a
    /// pasted screenshot arrives as the PNG it was. Anything else (HEIC from
    /// Photos, TIFF, an app's private type) is decoded and re-encoded.
    private static let portable: [UTType] = [.png, .jpeg, .gif, .webP]

    /// Reads the first pasteboard item. Called only from the user's own Paste
    /// command, which is why iOS shows no paste prompt and why Ferry still
    /// never reads the pasteboard on its own.
    static func read(from pasteboard: UIPasteboard) -> PastedImage? {
        guard pasteboard.hasImages else { return nil }
        let types = pasteboard.types.compactMap { UTType($0) }
        if let type = types.first(where: { candidate in portable.contains { candidate.conforms(to: $0) } }),
           let data = pasteboard.data(forPasteboardType: type.identifier),
           let mediaType = type.preferredMIMEType, let ext = type.preferredFilenameExtension {
            return PastedImage(data: data, mediaType: mediaType, filenameExtension: ext)
        }
        guard let image = pasteboard.image else { return nil }
        let hasAlpha = switch image.cgImage?.alphaInfo {
        case .none?, .noneSkipFirst?, .noneSkipLast?: false
        default: true
        }
        if hasAlpha, let data = image.pngData() {
            return PastedImage(data: data, mediaType: "image/png", filenameExtension: "png")
        }
        guard let data = image.jpegData(compressionQuality: 0.9) else { return nil }
        return PastedImage(data: data, mediaType: "image/jpeg", filenameExtension: "jpg")
    }
}

/// The composer's message field. SwiftUI's `TextField` offers Paste only when
/// the pasteboard holds text, so a copied screenshot could never reach Ferry.
/// This text view also accepts an image and hands it to `onPasteImage`; text
/// pastes as before.
struct MessageInput: UIViewRepresentable {
    @Binding var text: String
    @Binding var isFocused: Bool
    var isEnabled: Bool
    var maximumLines = 5
    var onPasteImage: (PastedImage) -> Void

    func makeUIView(context: Context) -> PastingTextView {
        let view = PastingTextView()
        view.delegate = context.coordinator
        view.font = .preferredFont(forTextStyle: .body)
        view.adjustsFontForContentSizeCategory = true
        view.backgroundColor = .clear
        view.textContainerInset = UIEdgeInsets(top: 10, left: 4, bottom: 10, right: 4)
        view.textContainer.lineFragmentPadding = 0
        view.isScrollEnabled = false
        view.setContentCompressionResistancePriority(.defaultLow, for: .horizontal)
        view.accessibilityIdentifier = "message-input"
        view.accessibilityLabel = "Message"
        view.placeholder.text = "Message"
        return view
    }

    func updateUIView(_ view: PastingTextView, context: Context) {
        context.coordinator.parent = self
        view.onPasteImage = onPasteImage
        if view.text != text { view.text = text }
        // Touched only on a change: restyling every update would rewrite the
        // attributes of text still being composed by a CJK input method.
        if view.isEditable != isEnabled {
            view.isEditable = isEnabled
            view.textColor = isEnabled ? .label : .secondaryLabel
            if isEnabled { view.accessibilityTraits.remove(.notEnabled) } else { view.accessibilityTraits.insert(.notEnabled) }
        }
        view.placeholder.isHidden = !text.isEmpty
        // Focus is applied after the current update so a field that has just
        // become editable can take the keyboard.
        if isFocused != view.isFirstResponder {
            DispatchQueue.main.async {
                if isFocused { view.becomeFirstResponder() } else { view.resignFirstResponder() }
            }
        }
    }

    func sizeThatFits(_ proposal: ProposedViewSize, uiView view: PastingTextView, context: Context) -> CGSize? {
        guard let width = proposal.width, width.isFinite, width > 0 else { return nil }
        let fitting = view.sizeThatFits(CGSize(width: width, height: .greatestFiniteMagnitude)).height
        let lineHeight = view.font?.lineHeight ?? UIFont.preferredFont(forTextStyle: .body).lineHeight
        let limit = (lineHeight * CGFloat(maximumLines) + view.textContainerInset.top + view.textContainerInset.bottom).rounded(.up)
        view.isScrollEnabled = fitting > limit
        return CGSize(width: width, height: min(fitting, limit))
    }

    func makeCoordinator() -> Coordinator { Coordinator(parent: self) }

    final class Coordinator: NSObject, UITextViewDelegate {
        var parent: MessageInput
        init(parent: MessageInput) { self.parent = parent }

        func textViewDidChange(_ textView: UITextView) {
            parent.text = textView.text
        }

        func textViewDidBeginEditing(_ textView: UITextView) {
            if !parent.isFocused { parent.isFocused = true }
        }

        func textViewDidEndEditing(_ textView: UITextView) {
            if parent.isFocused { parent.isFocused = false }
        }
    }
}

final class PastingTextView: UITextView {
    var onPasteImage: ((PastedImage) -> Void)?
    let placeholder = UILabel()

    override init(frame: CGRect, textContainer: NSTextContainer?) {
        super.init(frame: frame, textContainer: textContainer)
        placeholder.font = .preferredFont(forTextStyle: .body)
        placeholder.adjustsFontForContentSizeCategory = true
        placeholder.textColor = .placeholderText
        placeholder.isAccessibilityElement = false
        placeholder.translatesAutoresizingMaskIntoConstraints = false
        addSubview(placeholder)
        NSLayoutConstraint.activate([
            placeholder.leadingAnchor.constraint(equalTo: frameLayoutGuide.leadingAnchor, constant: 4),
            placeholder.topAnchor.constraint(equalTo: frameLayoutGuide.topAnchor, constant: 10),
        ])
    }

    required init?(coder: NSCoder) { fatalError("init(coder:) is not used") }

    override func canPerformAction(_ action: Selector, withSender sender: Any?) -> Bool {
        if action == #selector(paste(_:)), isEditable, UIPasteboard.general.hasImages { return true }
        return super.canPerformAction(action, withSender: sender)
    }

    override func paste(_ sender: Any?) {
        if let image = PastedImage.read(from: .general) {
            onPasteImage?(image)
        } else {
            super.paste(sender)
        }
    }
}
