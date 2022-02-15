
![logo](https://user-images.githubusercontent.com/43274508/154163236-a2c9679f-9f27-44a4-8f63-dc4774491a62.png)

This is the official CLI tool to operate with Getting Things Done Framework.

## How to use GTDF-CLI


#### 1. Create a new GTDF project

Create a new GTDF project in the linked version.

```bash
gtd new my_project
```



#### 2. Create a new view

Create a new directory with the typescript view inside. 

```
gtd view my_view 
```



#### 3. Create a new component

If a view name is passed as argument, the component will belong to that view directory, otherwise, the component will be global. 

```bash
gtd component my_component [my_view]
```



#### 4. Create a new service

Create a new typescript data service

```bash
gtd service my_service
```



#### 5. Create a new test

Create a new test. You can pass the test type, default value will be unit testing (u). 

```bash
gtd test my_test [u|i]
```
