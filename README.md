# shared
Library code shared among our repositories

## vin utils
Use the VIN type for strings that are VIN's and it gets some basic decoding functionality (eg. year, or tesla model).

## config loader
Convention over configuration library for typed configurations that can read from yaml file or env vars. Provides an opinionated
way of dealing with configurations: for local development we want to use yaml config files, and on higher level environments 
we want settings to come from environment variables (we use kubernetes). Supported types are: string, int, int64, bool. 
