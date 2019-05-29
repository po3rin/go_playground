FROM gradle:5.4.1-jdk11

COPY --chown=gradle:gradle . /home/gradle/src
WORKDIR /home/gradle/src/rescore
RUN gradle build
