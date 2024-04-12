FROM bitnami/node:20.12.2

RUN mkdir -p /home/apps/der-api-web

COPY . /home/apps/der-api-web

WORKDIR /home/apps/der-api-web

RUN npm install

EXPOSE 3000

ENTRYPOINT ["node", "index.js"]
