#FROM ubuntu:20.04
FROM scratch
#将打包好的linux代码映射到dokcer的app文件下
COPY main /app/pxx-server
#根目录是 /app
WORKDIR /app
#启动webook
CMD ["/app/pxx-server"]