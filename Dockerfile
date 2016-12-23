FROM scratch
MAINTAINER Bill Glover <billglover@gmail.com>
ADD load-sink /load-sink
EXPOSE 8080 8081
ENTRYPOINT ["/load-sink"]
