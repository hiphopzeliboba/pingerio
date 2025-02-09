# Используем Node.js для сборки
FROM node:18-alpine as build

WORKDIR /app
COPY . .
RUN npm install && npm run build

# Используем Nginx для раздачи статики
FROM nginx:alpine

COPY --from=build /app/build /usr/share/nginx/html

# Минимальная конфигурация для SPA
RUN printf "server {\n\
    listen 80;\n\
    location / {\n\
        root /usr/share/nginx/html;\n\
        index index.html;\n\
        try_files \$uri /index.html;\n\
    }\n\
}" > /etc/nginx/conf.d/default.conf

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
