DIRECTORY STRUCTURE
-------------------

      config/         yaml 文件结构体
      init/           项目初始化操作
      library/        扩展库
      models/         model层
      server/         app service层
      session/        websocket session list
      share/          进程全局变量
      static/         静态文件 http
      test/           测试
      vendor/         供应

Golang 1.9 版本 ！！！
------------
~~~
编译说明：
项目更目录 make 编译
或
go install 【编译后在golang bin目录下】
或
go build -o 可执行二进制文件名称 【编译后在当前目录】

启动服务
./编译后的二进制文件 -f yaml配置文件
./bnw_udp -f config.yaml
~~~
