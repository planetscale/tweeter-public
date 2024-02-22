# frozen_string_literal: true

source "https://rubygems.org"
git_source(:github) { |repo| "https://github.com/#{repo}.git" }

gem "rails", "~> 7.0.4"
gem "sprockets-rails"
gem "mysql2"
gem "active_storage_validations"
gem "bcrypt"
gem "bootsnap", require: false
gem "image_processing"
gem "inline_svg"
gem "puma"
gem "tzinfo-data", platforms: [:mingw, :mswin, :x64_mingw, :jruby]
gem "webpacker"
gem "will_paginate-bootstrap-style"
gem "will_paginate"
gem "dotenv"

# Use Redis adapter to run Action Cable in production
# gem 'redis', '~> 4.0'
# Build JSON APIs with ease. Read more: https://github.com/rails/jbuilder
# gem 'jbuilder', '~> 2.5'

group :development, :test do
  gem "capybara"
  gem "debug", platforms: %i[ mri mingw x64_mingw ]
  gem "factory_bot_rails"
  gem "faker"
  gem "net-smtp", require: false
  gem "rspec_junit_formatter"
  gem "rspec-rails"
  gem "selenium-webdriver"
  gem "simplecov", require: false
  gem "planetscale_rails"
end

group :development do
  gem "bullet"
  gem "listen"
  gem "rack-mini-profiler"
  gem "web-console"
end

group :rubocop do
  gem "code-scanning-rubocop", require: false
  gem "rubocop-rails_config"
  gem "rubocop"
end

gem "dockerfile-rails", ">= 1.5", :group => :development
