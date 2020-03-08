FROM busybox:1.31

COPY ./goechomariadb /home/
EXPOSE 8000

CMD ["/home/goechomariadb"]
