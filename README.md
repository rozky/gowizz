# WizzAir Flights API

Existing scrappers:

https://github.com/alexeiTruhin/wizzpricehistory/blob/master/main/wizz-api.js

https://gist.github.com/YoanaTodorova/4eb820cdf3fe7b1a3d020bf145696f5d

API Errors:
InvalidFromDate -> if the from date is in the past
InvalidMarket -> 
    - if there is no connection between airports provided in the search request
    - if airport code is not valid
InvalidTimeDateRange -> if time range (from, to) is bigger than 40 days (or something very close to it)    


## Serverless

### Working without admin permissions 
```
npm install serverless --save-dev
npx serverless deploy
```