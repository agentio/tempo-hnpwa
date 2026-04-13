all:
	cd ui; pnpm install; pnpm build
	cp ui/dist/assets/index.js internal/assets/js
	cp ui/dist/assets/index.css internal/assets/css
	cp ui/public/icon-512x512.png internal/assets/img
	go install ./...
