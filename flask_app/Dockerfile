FROM python:3.9.7-slim-buster

WORKDIR /app
ENV FLASK_APP=app

COPY ./src/requirements.txt ./

# for installing uWSGI and mysql-client
RUN apt-get update && \
    apt-get -y install gcc libmariadb-dev

RUN pip install --upgrade pip
RUN pip install -r requirements.txt