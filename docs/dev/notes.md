# 笔记

本地开发环境搭建，参考 [local_development.md](local_development.md)

## 配置文件加载流程

优先从  `custom/conf/app.ini` 加载配置， 如果该文件不存在，再从 `conf/app.ini` 加载。

相关代码如下：

```
// CustomDir returns the absolute path of the custom directory that contains local overrides.
// It reads the value of environment variable GOGS_CUSTOM. When not set, it uses the work
// directory returned by WorkDir function.
func CustomDir() string {
	customDirOnce.Do(func() {
		customDir = os.Getenv("GOGS_CUSTOM")
		if customDir != "" {
			return
		}

		customDir = filepath.Join(WorkDir(), "custom")
	})

	return customDir
}
```

```
func Init(customConf string) error {
	data, err := conf.Files.ReadFile("app.ini")
	if err != nil {
		return errors.Wrap(err, `read default "app.ini"`)
	}

	File, err = ini.LoadSources(ini.LoadOptions{
		IgnoreInlineComment: true,
	}, data)
	if err != nil {
		return errors.Wrap(err, `parse "app.ini"`)
	}
	File.NameMapper = ini.SnackCase

	if customConf == "" {
		customConf = filepath.Join(CustomDir(), "conf", "app.ini")
	} else {
		customConf, err = filepath.Abs(customConf)
		if err != nil {
			return errors.Wrap(err, "get absolute path")
		}
	}
	CustomConf = customConf

	if osutil.IsFile(customConf) {
		if err = File.Append(customConf); err != nil {
			return errors.Wrapf(err, "append %q", customConf)
		}
	} else {
		log.Warn("Custom config %q not found. Ignore this warning if you're running for the first time", customConf)
	}

	if err = File.Section(ini.DefaultSection).MapTo(&App); err != nil {
		return errors.Wrap(err, "mapping default section")
	}
    // 省略 ......
}
```

## /:username/:reponame 路由


## 国际化语言切换


