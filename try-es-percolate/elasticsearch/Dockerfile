FROM docker.elastic.co/elasticsearch/elasticsearch:6.6.1

# x-pack はインストールしない
#```
#ERROR: this distribution of Elasticsearch contains X-Pack by default
#ERROR: Service 'elasticsearch' failed to build: The command '/bin/sh -c elasticsearch-plugin install --batch x-pack' returned a non-zero code: 78
#```
#RUN elasticsearch-plugin install --batch x-pack

# # kuromojiをインストール
# RUN elasticsearch-plugin install analysis-kuromoji
# RUN elasticsearch-plugin install --batch org.codelibs:elasticsearch-analysis-kuromoji-neologd:6.6.0
