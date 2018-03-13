FROM golang

# if left blank app will run with dev settings
# to build production image run:
# $ docker build ./api --build-args app_env=production
ARG app_env
ENV APP_ENV $app_env

# it is okay to leave user/GoDoRP as long as you do not want to share code with other libraries
COPY . /go/src/github.com/yusufhacktiv8/pediatric-eye-clinic
WORKDIR /go/src/github.com/yusufhacktiv8/pediatric-eye-clinic

RUN echo $PATH

RUN go get ./...

RUN go build

# if dev setting will use pilu/fresh for code reloading via docker-compose volume sharing with local machine
# if production setting will build binary
CMD if [ ${APP_ENV} = production ]; \
	then \
	pediatric-eye-clinic; \
	elif [ ${APP_ENV} = test ]; \
        then \
        go test -v ./...; \
        else \
	go get github.com/pilu/fresh && \
	fresh; \
	fi

EXPOSE 8080
