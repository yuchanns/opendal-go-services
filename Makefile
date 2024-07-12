.PHONY: build-aliyun-drive build-memory

build-aliyun-drive:
	cd opendal/bindings/c && \
	cargo build --release && \
	cp ./target/release/libopendal_c.so ../../../aliyun_drive/libopendal_c.so && \
	cd - && \
	cd aliyun_drive && \
	zstd -22 libopendal_c.so && \
	rm libopendal_c.so && \
	mv libopendal_c.so.zst libopendal_c.so.arm64.zst

build-memory:
	cd opendal/bindings/c && \
	cargo build --release && \
	cp ./target/release/libopendal_c.so ../../../memory/libopendal_c.so && \
	cd - && \
	cd memory && \
	zstd -22 libopendal_c.so && \
	rm libopendal_c.so && \
	mv libopendal_c.so.zst libopendal_c.so.arm64.zst
