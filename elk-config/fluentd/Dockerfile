# Use Ruby 3 as the base image
FROM ruby:3.0

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

# Expose Fluentd ports
EXPOSE 24224 24224/udp

# Run Fluentd with the custom configuration file
CMD ["fluentd", "-c", "/fluentd/etc/fluent.conf"]
