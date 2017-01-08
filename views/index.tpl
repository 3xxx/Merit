<!-- 首页界面 -->
<!DOCTYPE html>
<html>
<head>
<title>Merit价值管理系统</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  

  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <!-- <meta name="author" content="Jophy" /> -->
  <!-- <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>  -->
  <!-- <link rel="stylesheet" href="/static/css/style.css"> -->
  <!--[if lte IE 9]>兼容ie的方面：先引用bootstrapcss，再引用js-->
 
<script src="/static/js/html5.js"></script>
<script src="/static/js/respond.min.js"></script>
<script src="/static/js/html5shiv.min.js"></script>
    <!--[if lt IE 9]>
    <script src="https://www.novamind.com/wp-content/themes/novamind/js/html5.js"></script>
    <script src="https://oss.maxcdn.com/html5shiv/3.7.2/html5shiv.min.js"></script>
    <script src="https://oss.maxcdn.com/respond/1.4.2/respond.min.js"></script>
    <![endif]-->

<!-- <link rel="stylesheet" id="genericons-css" href="/static/NovaMind/genericons.css" type="text/css" media="all"> -->
<link rel="stylesheet" id="novamind-style-css" href="/static/NovaMind/style.css" type="text/css" media="all"><!--这个是下面的样式-->


<link rel="stylesheet" id="Novamindtwelve-style-css" href="/static/NovaMind/style3.css" type="text/css" media="all"><!--这个是上面的样式-->
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
<style type="text/css">
div #footer {
color: #E6E6FA;
}
a {
color: #F5DEB3;
}
</style>
</head>

<body>
  <!--<div class="col-lg-12">
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

    <form class="form-inline" method="post" action="/import_xls_catalog" enctype="multipart/form-data">
      <div class="form-group">
        <label>选择excel</label>
        <input type="file" class="form-control" name="catalog" id="catalog"></div>
      <button type="submit" class="btn btn-default">提交</button>
    </form>

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
          <span class="glyphicon glyphicon-user"></span>
          查看价值结构
        </button>
      </a>
    </div>
    <div class="form-group">
      <a href="/getperson">
        <button type="button" class="btn btn-primary">
          <span class="glyphicon glyphicon-user"></span>
          价值排序
        </button>
      </a>
    </div>
    <div class="form-group">
      <a href="/user">
        <button type="button" class="btn btn-primary">
          <span class="glyphicon glyphicon-user"></span>
          查看个人
        </button>
      </a>
    </div>
  </div>-->


<div class="container-fluid blue-bg-rpt">
  <div class="row">
    <h2 class="how-it-works">How can Merit help me?</h2>
    <div class="col-xs-6 col-sm-3 col-md-3 col-lg-3">
      <div class="hlp-box presentation home-how-box" id="home-presentation" style="cursor: pointer;">
        <h3>成果登记</h3>
        <p>
          设计人员轻松进行成果登记；管理人员轻松对成果进行分析。
        </p>
      </div>
    </div>

    <div class="col-xs-6 col-sm-3 col-md-3 col-lg-3">
      <div class="hlp-box visual-pl home-how-box" id="visualpl" style="cursor: pointer;">
        <h3>价值管理</h3>
        <p>基于自我维护的档案管理理念。<br/>
        发现价值、展示价值<br/>
        积累价值、提升价值
        </p>
      </div>
    </div>
    <div class="col-xs-6 col-sm-3 col-md-3 col-lg-3">
      <div class="hlp-box studying home-how-box" id="home-studying" style="cursor: pointer;">
        <h3>奖金分配</h3>
        <p>根据项目产值、成果排名和个人价值进行分配，即双轨制奖金分配方案。</p>
      </div>
    </div>
    <div class="col-xs-6 col-sm-3 col-md-3 col-lg-3">
      <div class="hlp-box pitching home-how-box" id="home-pitching" style="cursor: pointer;">
        <h3>EngineerCMS</h3>
        <p>
          基于工程师个人电脑的在线知识管理系统，优雅地对个人项目和设计成果进行网络化管理。
        </p>
      </div>
    </div>

  </div>
<br/>
<br/>
<br/>
  <div class="col-md-12">
    <div class="col-480-12 col-xs-4 col-sm-4 col-md-4">
      <div class=" broad-feature-12">
        <div class="broad-feature feature1-box">
          <div class="col-480-12">
            <h3 class="col-480-12">高效</h3>
            <p class="col-480-12 broad-box-p">
              让工程师专心于高级技术应用，减少琐事.
            </p>

          </div>
        </div>
      </div>
    </div>
    <div class="col-480-12 col-xs-4 col-sm-4 col-md-4">
      <div class=" broad-feature-12">
        <div class="broad-feature feature2-box">
          <div class="col-480-12">
            <h3 class="col-480-12">弹性</h3>
            <p class="col-480-12 broad-box-p">
              自由地进行成果查看、统计分析.
            </p>

          </div>
        </div>
      </div>
    </div>
    <div class="col-480-12 col-xs-4 col-sm-4 col-md-4">
      <div class="col-480-12 broad-feature-12 ">
        <div class=" broad-feature feature3-box">
          <div class="col-480-12">
            <h3 class="col-480-12">优雅</h3>
            <p class="col-480-12 broad-box-p">
              利用网络化、自动化管理，提高品质.
            </p>

          </div>
        </div>
      </div>
    </div>
  </div>

<div id="footer">
  <div class="col-lg-12">
    <br>
    <hr/>
  </div>

  <div class="col-lg-6">
    <h4>Copyright © 2016 Merit</h4>
    <p>
      网站由 <i class="user icon"></i>
      <a target="_blank" href="https://github.com/3xxx">@3xxx</a>
      建设，并由
      <a target="_blank" href="http://golang.org">golang</a>
      和
      <a target="_blank" href="http://beego.me">beego</a>
      提供动力。
    </p>

    <p>
      请给 <i class="glyphicon glyphicon-envelope"></i>
      <a class="email" href="mailto:qin.xc@gpdiwe.com">我</a>
      发送反馈信息或提交
      <i class="tasks icon"></i>
      <a target="_blank" href="https://github.com/3xxx/merit/issues">网站问题</a>
      。
    </p>
  </div>
  <div class="col-lg-6">
    <h4 >更多项目</h4>
    <div >
      <p>
        <a href="https://github.com/3xxx/hydrows">HydroWS水利供水管线设计工具</a>
      </p>
      <p>
        <a href="https://github.com/3xxx/engineercms">EngineerCMS工程师知识管理系统</a>
      </p>
    </div>
  </div>
</div>

</div>


  <!--<div class="gm-section-5">
    <div class="col-md-12">
      <div class="col-480-12 col-xs-4 col-sm-4 col-md-4">
        <div class=" broad-feature-12">
          <div class="broad-feature feature1-box">
            <div class="col-480-12">
              <h3 class="col-480-12">Merit</h3>
              <p class="col-480-12 broad-box-p">
                价值管理系统.
              </p>

            </div>
          </div>
        </div>
      </div>
      <div class="col-480-12 col-xs-4 col-sm-4 col-md-4">
        <div class=" broad-feature-12">
          <div class="broad-feature feature2-box">
            <div class="col-480-12">
              <h3 class="col-480-12">High Level Overview</h3>
              <p class="col-480-12 broad-box-p">
                成果登记系统.
              </p>

            </div>
          </div>
        </div>
      </div>
      <div class="col-480-12 col-xs-4 col-sm-4 col-md-4">
        <div class="col-480-12 broad-feature-12">
          <div class=" broad-feature feature3-box">
            <div class="col-480-12">
              <h3 class="col-480-12">
                MS Project Import &amp; Export
              </h3>
              <p class="col-480-12 broad-box-p">
                HydroCMS水利设计管理系统.
              </p>

            </div>
          </div>
        </div>
      </div>
    </div>
  </div>-->

<script>
// $('#getstarted').click(function () {
        
//      location.href='/download/?id='+$('#download_app').val();
//      });


$('#home-presentation').css({'cursor': 'pointer'});
    
$('#home-presentation').click(function () {
    
location.href='/achievement';
  
});

$('#visualpl').css({'cursor': 'pointer'});
    
$('#visualpl').click(function () {
    
location.href='/merit';
  
});

$('#home-studying').css({'cursor': 'pointer'});
    
$('#home-studying').click(function () {
    
location.href='/dollars';
  
});

$('#home-pitching').css({'cursor': 'pointer'});
    
$('#home-pitching').click(function () {
    
location.href='http://192.168.9.13';
  
});





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
<span class="glyphicon glyphicon-user"></span>
User
</button>
-->