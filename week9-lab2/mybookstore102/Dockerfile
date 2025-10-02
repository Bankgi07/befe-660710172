FROM node:18-alpine AS builder

WORKDIR /app

COPY package*.json ./

# ติดตั้งให้รวม devDependencies ด้วย
ENV NODE_ENV=development
RUN npm install

COPY . .

# แก้ permission ก่อน build
RUN chmod -R 755 node_modules/.bin

RUN npm run build

FROM nginx:alpine

COPY --from=builder /app/build /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
