FROM ubuntu

MAINTAINER kuromike0629 <kuromike0629@gmail.com>

#Install required packages
RUN apt-get update && apt-get install -y git make g++

#Get Gikou from github.com
WORKDIR /tmp
RUN git clone http://github.com/gikou-official/Gikou.git

#Make Gikou
WORKDIR /tmp/Gikou
RUN make release

