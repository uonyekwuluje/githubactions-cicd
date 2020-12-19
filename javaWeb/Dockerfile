FROM tomcat:8.0
LABEL maintainer="uionyekwuluje@gmail.com"
ARG JAR_FILE=target/helloworld.war
RUN rm -rf /usr/local/tomcat/webapps/* 
COPY ${JAR_FILE} /usr/local/tomcat/webapps/ROOT.war
CMD ["catalina.sh","run"]
