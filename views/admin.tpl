<!-- 首页界面 -->
<!DOCTYPE html>
<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script> 
 <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
</head>

<body>
<div class="col-lg-12">
<div class="form-group">
  <form method="post" action="/importjson" enctype="multipart/form-data">
    <div class="input-group">
      <label>
        选择json数据文件：
        <input type="file" name="json" id="json" />
      </label>
      <br/>
    </div>
    <button type="submit" class="btn btn-default" >导入结构数据</button>
  </form>
</div>

<div class="form-group">
  <button type="button" class="btn btn-default" id="import">初始化评测结构数据</button>
</div>

<div class="form-group">
{{if .IsLogin}}
    <a href="/login?exit=true">
    <button type="button" class="btn btn-primary">
      <span class="glyphicon glyphicon-user"></span>
      管理员退出
    </button>
    </a>
        {{else}}
    <a href="/login?url=/admin">
    <button type="button" class="btn btn-primary">
      <span class="glyphicon glyphicon-user"></span>
      管理员登录
    </button>
    </a>
{{end}}
</div>

  <div class="form-group">
  <a href="/json">
    <button type="button" class="btn btn-primary">
      <!--  btn-lg -->
      <span class="glyphicon glyphicon-user"></span>
      查看价值结构
    </button>
  </a>
  </div>
  <div class="form-group">
    <a href="/getperson">
    <button type="button" class="btn btn-primary">
      <!--  btn-lg -->
      <span class="glyphicon glyphicon-user"></span>
      价值排序
    </button>
  </a>  
  </div>
  <div class="form-group">
    <a href="/user">
    <button type="button" class="btn btn-primary">
      <!--  style="text-shadow: black 5px 3px 3px;" btn-lg -->
      <span class="glyphicon glyphicon-user"></span>
      查看个人
    </button>
  </a>
</div>  
</div>
  <script>
$(document).ready(function(){
$("#import").click(function(){//这里应该用button的id来区分按钮的哪一个,因为本页有好几个button
            $.ajax({
                type:"POST",
                url:"/importjson",
                success:function(data){//数据提交成功时返回数据
                    alert("导入成功！")
                }
            });
            return true;//这里true和false结果都一样。不刷新页面的意思？
 });
});
  </script>
</body>
</html>
<!-- <button type="button" class="btn btn-primary btn-lg" style="color: rgb(212, 106, 64);">
<span class="glyphicon glyphicon-user"></span> User
</button> -->