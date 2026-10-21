---
scope: Ferry's Server, Web app, iOS and Android apps, HTTP contract and data backup/restore utility.
method: Static tracing of tracked source, existing specifications and selected test cases.
covered:
- description: Server admission, device management, timeline, file storage and download-link code.
  paths:
  - internal/ferry/
- description: Browser client behavior and entry surface.
  paths:
  - internal/webui/
- description: Native client connection, composer, timeline, image and export code.
  paths:
  - ios/Ferry/Ferry/
  - android/app/src/main/java/
- description: Server startup and the data backup and restore utility.
  paths:
  - cmd/ferry/
  - scripts/ferry-data.sh
  - docker-entrypoint.sh
  - compose.yaml
exclusions:
- description: Build, packaging, CI, artwork and release automation; these are contributor delivery mechanisms.
  paths:
  - .github/
  - Dockerfile
  - compose.build.yaml
  - scripts/check-repo.sh
  - scripts/image-smoke.sh
  - AppIcon.appiconset/
  - Artwork/
  - android/gradle/
  - android/build.gradle.kts
  - android/settings.gradle.kts
  - android/app/build.gradle.kts
  - ios/Ferry/Ferry.xcodeproj/
- description: Native app resources and styling assets; these are design rather than product behavior.
  paths:
  - android/app/src/main/res/
  - ios/Ferry/Ferry/Assets.xcassets/
- description: Native app test suites.
  paths:
  - android/app/src/test/
  - ios/Ferry/FerryTests/
  - ios/Ferry/FerryUITests/
unmapped:
- description: Database schema migration code; the retained-history outcome is covered, the migration steps are not.
  paths:
  - internal/ferry/store.go
limitations: []
---

# Coverage
