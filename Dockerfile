FROM postgres:latest

# Set Environment Variables
ENV POSTGRES_DB=postgres
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=postgres

# Expose the PostreSQL Port
EXPOSE 5432