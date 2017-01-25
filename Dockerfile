FROM scratch
MAINTAINER Bill Glover <billglover@gmail.com>
ADD load-sink /load-sink
EXPOSE 8080
ENTRYPOINT ["/load-sink", "-delay=500", "-jitter=100", "-size=50", "-variance=10"]
