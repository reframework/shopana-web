FROM node:18-alpine as build
WORKDIR /app

ARG NEXT_PUBLIC_WEBAPP_ROLLBAR_TOKEN
ARG NEXT_PUBLIC_WEBAPP_GOOGLE_ANALYTICS_ID
ARG NEXT_PUBLIC_WEBAPP_GRAPHQL_URL

ENV NEXT_PUBLIC_WEBAPP_ROLLBAR_TOKEN=$NEXT_PUBLIC_WEBAPP_ROLLBAR_TOKEN
ENV NEXT_PUBLIC_WEBAPP_GA_TOKEN=$NEXT_PUBLIC_WEBAPP_GA_TOKEN
ENV NEXT_PUBLIC_WEBAPP_GRAPHQL_URL=$NEXT_PUBLIC_WEBAPPGRAPHQL_URL

COPY package.json yarn.lock ./
RUN yarn install --immutable --frozen-lockfile

COPY .eslintrc.json next.config.js tsconfig.json ./
COPY ./src ./src
COPY ./public ./public
RUN yarn build

FROM node:18-alpine as production
WORKDIR /app
COPY --from=build /app/package.json /app/next.config.js ./
COPY --from=build /app/.next ./.next
COPY --from=build /app/public ./public
COPY --from=build /app/node_modules ./node_modules

# Disable Next.js telemetry
ENV NEXT_TELEMETRY_DISABLED 1
CMD ["yarn", "start"]
