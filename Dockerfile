FROM scratch
MAINTAINER Bill Glover <billglover@gmail.com>
ADD ./load-sink /load-sink
ENTRYPOINT ["/load-sink"]
