FROM node:14-bullseye-slim
ENV NODE_ENV=production
WORKDIR /app
COPY "package.json" "./"
RUN npm install --production
COPY . /app
EXPOSE 8080
CMD [ "node", "server.js" ]
