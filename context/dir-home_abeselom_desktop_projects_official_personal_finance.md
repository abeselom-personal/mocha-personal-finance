# Source Code Context

Generated on: 2025-06-09T18:01:31Z

## Repository Overview
- Total Files: 5
- Total Size: 195 bytes

## Directory Structure
```
Makefile
Readme.md
context/
  images/
front-end/
  mobile/
  web/
services/
  auth/
    cmd/
    config/
    controller/
    db/
    model/
    routes/
    service/
  gateway/
    Dockerfile
    certs/
    nginx.conf
  notification/
    cmd/
    config/
    controller/
    db/
    model/
    routes/
    service/
  receipt-scraper/
    config/
    controller/
    db/
    main.py
    model/
    queue/
    scraper/
  regex/
    cmd/
    config/
    controller/
    db/
    model/
    routes/
    service/
  shared/
    constants/
    contracts/
    middleware/
    proto/
    utils/
  sms-parser/
    cmd/
    config/
    controller/
    db/
    model/
    routes/
    service/
  sync/
    cmd/
    config/
    controller/
    db/
    model/
    routes/
    service/

```

## File Contents


### File: Makefile

```

.PHONY: build up down restart logs

build:
	docker-compose build

up:
	docker-compose up -d

down:
	docker-compose down

restart:
	make down && make up

logs:
	docker-compose logs -f --tail=100

```





### File: Readme.md

```markdown

```





### File: services/gateway/Dockerfile

```

```





### File: services/gateway/nginx.conf

```

```





### File: services/receipt-scraper/main.py

```python

```




