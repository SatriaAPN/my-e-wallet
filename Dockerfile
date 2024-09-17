# Use the official PostgreSQL image as the base image
FROM postgres:15

# Set environment variables for PostgreSQL user, password, and database
ENV POSTGRES_USER=myuser
ENV POSTGRES_PASSWORD=mypassword
ENV POSTGRES_DB=mydatabase

# Copy initialization SQL script to the directory that PostgreSQL uses for initialization
COPY ./init.sql /docker-entrypoint-initdb.d/

# Expose the PostgreSQL default port
EXPOSE 5432
