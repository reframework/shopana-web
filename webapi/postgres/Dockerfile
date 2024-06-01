FROM postgres:16

COPY ./init.sql /docker-entrypoint-initdb.d/

# Install necessary packages
RUN apt-get update && apt-get install -y \
    build-essential \
    postgresql-server-dev-all \
    git

# Clone the extension repository
RUN git clone https://github.com/fboulnois/pg_uuidv7.git

# Build and install the extension
RUN cd pg_uuidv7 && make && make install
