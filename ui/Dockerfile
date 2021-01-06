FROM node:15.4.0 AS compile-image
WORKDIR /opt/ng
COPY package.json ./
RUN npm install
ENV PATH="./node_modules/.bin:$PATH" 
COPY . ./
RUN ng build --prod
# Switch if you want to build a Docker image used in k8s (local + NodePort)
# See environment.k8s.ts
#RUN ng build --configuration k8s

FROM nginx
COPY nginx.conf /etc/nginx/nginx.conf
COPY --from=compile-image /opt/ng/dist/ui /usr/share/nginx/html