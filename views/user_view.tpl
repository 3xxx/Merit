<!-- 用户修改自己的密码 -->
<!DOCTYPE html>
<html>
<head>
<title>Merit价值管理系统</title>
  <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
</head>
<body>

<div class="col-lg-12">
<div><h2>鼠标离开输入框即修改完成</h2></div>
<table class="table table-striped">
  <thead>
      <tr>
      <th><span style="cursor: pointer">Username</span></th>
      <th><span style="cursor: pointer">Password</span></th>
      <th><span style="cursor: pointer">Nickname</span></th>
      <th><span style="cursor: pointer">Email</span></th>
      <th><span style="cursor: pointer">分院</span></th>
      <th><span style="cursor: pointer">科室</span></th>
      <th><span style="cursor: pointer">Lastlogintime</span></th>
      <th><span style="cursor: pointer">Createtime</span></th>
      <th><span style="cursor: pointer">权限Role</span></th>
    </tr>
  </thead>

<tbody>

    <tr><!--tr表格的行，td定义一个单元格，<th> 标签定义表格内的表头单元格-->
      <td>{{.User.Username}}</td>
      <td><input type="password" class="form-control" id="input" name="password"  size='5'/></td>
      <td><input type="text" class="form-control" id="input" name="nickname" value="{{.User.Nickname}}" size='6'/></td>
      <td><input type="text" class="form-control" id="input" name="email" value="{{.User.Email}}" size='20'/></td>
      <td>{{.User.Department}}</td>
      <td>{{.User.Secoffice}}</td>
      <td>{{dateformat .User.Lastlogintime "2006-01-02 T 15:04:05"}}</td>
      <td>{{dateformat .User.Createtime "2006-01-02 T 15:04:05"}}</td>
      <td>{{.User.Role}}</td>
    </tr>
 </tbody>   
 </table>
</div>

<script>
$(document).ready(function(){
  // var roletitle1=document.getElementsByName("roletitle");
  // $("#uname").focus(function(){获得焦点
     $("input").blur(function(){//其失去焦点
        var pwd=document.getElementsByName("password");
        var nickname=document.getElementsByName("nickname");
        var email=document.getElementsByName("email");
        // var roletitle2=document.getElementsByName("roletitle");
        // alert(pwd[0].value.length);//什么时候是逗号，什么时候是分号？
        if (pwd[0].value.length<1)
        {
          // alert("请输入密码。")
             $.ajax({
                type:"post",//这里是否一定要用post？？？
                url:"/user/UpdateUser",
                data: { userid:{{.User.Id}},username:{{.User.Username}},nickname: nickname[0].value,email: email[0].value},
                success:function(data,status){//数据提交成功时返回数据
                  // alert(data,status);
                 }
            });         
        }else{
            $.ajax({
                type:"post",//这里是否一定要用post？？？
                url:"/user/UpdateUser",
                data: { userid:{{.User.Id}},username:{{.User.Username}},password: pwd[0].value,nickname: nickname[0].value,email: email[0].value},
                success:function(data,status){//数据提交成功时返回数据
                  alert('success modified~');
                  // alert(data,status);
                 }
            });
       }     
 });
});



function checkInput(){
  var uname=document.getElementById("uname");
  if (uname.value.length==0){
    alert("请输入账号");
    return false;
  }
    var pwd=document.getElementById("pwd");
  if (pwd.value.length==0){
    alert("请输入密码");
    return false;
    }
// return true
}

</script>

</body>
</html>