FROM ruby:3.0

# Create a user for the container
RUN addgroup --system fluent && adduser --system --ingroup fluent fluent

# Install Fluentd and required plugins
RUN gem install fluentd -v 1.13.3 && \
    gem install fluent-plugin-elasticsearch -v 5.0.0 && \
    gem install elasticsearch-transport -v 7.13.3

# Create a directory for Fluentd configuration
RUN mkdir -p /fluentd/etc

# Set the default working directory
WORKDIR /fluentd

# Copy Fluentd configuration file
COPY /fluentd/fluent.conf /fluentd/etc/fluent.conf

# Set the user to run Fluentd
USER fluent

# Add a health check
HEALTHCHECK --interval=30s --timeout=10s --retries=3 CMD curl -f http://localhost:24224/api/plugins.json || exit 1

# Run Fluentd with the custom configuration file
CMD ["fluentd", "-c", "/fluentd/etc/fluent.conf"]
