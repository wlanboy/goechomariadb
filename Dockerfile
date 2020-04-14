FROM busybox:1.31

RUN mkdir -p /usr/local/share/busybox && echo "/bin/busybox sh" > /usr/local/share/busybox/sh && chmod +x /usr/local/share/busybox/sh
RUN addgroup -S kanban && adduser -S kanban -G kanban -s /usr/local/share/busybox/sh

COPY ./goechomariadb /home/kanban

USER kanban

EXPOSE 8000

CMD ["/home/kanban/goechomariadb"]
