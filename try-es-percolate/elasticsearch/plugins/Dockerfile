FROM gradle:jdk8-alpine

COPY --chown=gradle:gradle . /home/gradle/src
WORKDIR /home/gradle/src/rescore
RUN gradle build
