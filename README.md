# REPO

This repository contains how we, developers, use buildpacks behind the scene on Google Cloud Platform.
These products includes App Engine, Cloud Functions, and optional for Cloud Run.

App Engine and Cloud Functions uses buildpacks.

Google's buildpacks builder has two flavors, one for App Engine, the other for the remains.
According to Google, the "App Engine flavor"-ed builder, adds some App Engine specific metadata. They do not have big differences.

And the builder also have different "runtime images", Ubuntu 18 and Ubuntu 22, respects to what runtime you choosed in App Engine or Cloud Functions, the proper "run image" is choosed for you.
