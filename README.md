人大金仓Gorm驱动

### 依赖下载:
```shell
go get github.com/x-alchemist/gorm-kingbase-driver@v1.0.1
```

### 配置go.mod依赖:
``` go.mod 
require(
    github.com/x-alchemist/gorm-kingbase-driver v1.0.1
)
```

```go
import  github.com/x-alchemist/gorm-kingbase-driver

var db *gorm.DB
	var err error
	if db, err = gorm.Open(gorm-kingbase-driver.New(kingbaseConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	}); err != nil {
		return nil, err
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(k.MaxIdleConns)
		sqlDB.SetMaxOpenConns(k.MaxOpenConns)
		return db, nil
	}
```