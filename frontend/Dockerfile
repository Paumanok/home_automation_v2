FROM node:lts-alpine

# make the 'app' folder the current working directory
WORKDIR /usr/app

COPY ./ /usr/app

# install project dependencies
RUN npm install

# build app for production with minification
#RUN npm run build

#EXPOSE 5000
#CMD [ "http-server", "dist" ]
CMD ["npm", "run", "dev", "--", "--host"]
