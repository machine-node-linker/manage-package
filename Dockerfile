FROM scratch
ENTRYPOINT [ "/manage-package" ]
COPY manage-package /

ARG VERSION
ENV VERSION=${VERSION:dev}