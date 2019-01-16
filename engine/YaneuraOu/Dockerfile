FROM ubuntu

MAINTAINER kuromike0629 <kuromike0629@gmail.com>

ENV DEBIAN_FRONTEND=noninteractive

#Install required packages
RUN apt-get update && apt-get install -y make g++ mono-complete unzip wget
RUN apt-get install -y tzdata

#Get Gikou from github.com
RUN wget https://github.com/yaneurao/YaneuraOu/archive/v4.79engine2017.tar.gz
RUN tar xvzf v4.79engine2017.tar.gz

WORKDIR /YaneuraOu-4.79engine2017/source/

#Make YaneuraOu
RUN make sse2 COMPILER=g++

CMD ./YaneuraOu-by-gcc
