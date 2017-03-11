<!-- 后台主页面，其他为子页面-->
<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <title>EngineerCMS</title>

  <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <script src="/static/js/bootstrap-treeview.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-treeview.css"/>
  <!-- <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script> -->
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
  <!-- <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css"/> -->
  <!-- <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-editable.css"/> -->
  <!-- <script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script> -->
  <!-- <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script> -->
  <!-- <script type="text/javascript" src="/static/js/bootstrap-table-editable.min.js"></script> -->
  <!-- <script type="text/javascript" src="/static/js/bootstrap-editable.js"></script> -->
  <!-- <script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script> -->

  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css"/>
  <!-- <script src="/static/js/tableExport.js"></script> -->
</head>
<body>
<div class="col-lg-2">
  <div id="tree"></div>
</div>
<!-- 菜单顶部 -->
  <div class="col-lg-10">
    <div class="navbar navbar-top">
      <ul class="nav navbar-nav navbar-right">
        <li>
          <a href="/achievement">成果</a>
        </li>
        <li>
          <a href="/merit">价值</a>
        </li>
        <li>
          <a href="/dollars">奖金</a>
        </li>
        <li>
          <a href="/admin/user/detail">{{.Ip}}</a>
        </li>
        <li class="dropdown">
          <a href="#" class="dropdown-toggle" data-toggle="dropdown">Dropdown<b class="caret"></b></a>
          <ul class="dropdown-menu">
            <li><a href="https://github.com/3xxx" target="_blank">3xxx github</a></li>
            <li><a href="http://www.sina.com/xc-qin" target="_blank">Weibo</a></li>
            <li><a href="#">Something</a></li>
            <li class="divider"></li>
            <li><a href="http://blog.csdn.net/hotqin888" target="_blank">Blog</a></li>
          </ul>
        </li>
        <li>
          <a href="/admin/login/out">Log out</a>
        </li>
      </ul>
    </div>

    <div class="breadcrumbs">
      <ol class="breadcrumb" split="&gt;">
        <li>
          <a href="javascript:void(0)"> <i class="fa fa-home" aria-hidden="true"></i>
          后台
          </a>
        </li>
        <li>
          <a href="javascript:void(0)"> <i class="fa '. $parents['picon'] .' " aria-hidden="true"></i>
          系统设置
          </a>
        </li>
        <li>
          <a href="javascript:void(0)">
            <i class="fa '. $parents['icon'] .' " aria-hidden="true"></i>组织结构
          </a>
        </li>
      </ol>
    </div>
  </div>

  <script type="text/javascript">
    $(function () {
      var data = 
      [
        {
          text: "欢迎您~{{.Ip}}",
          icon: "fa fa-optin-monster",
          selectable: true,
          id: '010',
        },
        {
          text: "系统设置",
          icon: "fa fa-tachometer icon",
          // selectedIcon: "glyphicon glyphicon-stop",
          href: "#node-1",
          selectable: true,
          id: '01',
          selectable: false,
          // state: {
            // checked: true,
            // disabled: true,
            // expanded: true,
            // selected: true
          // },
          tags: ['available'],
          nodes: 
          [
            { 
              icon: "fa fa-cog",
              text: "基本设置",
              id: '011',
              nodeId: '011'
            },
            { 
              icon: "fa fa-align-right",
              text: "组织结构",
              id: '012',
              nodeId: '012'
            }, 
            { 
              icon: "fa fa-dollar",
              text: "定义价值",
              id: '013'
            }, 
            { 
              icon: "fa fa-cny",
              text: "科室价值",
              id: '014'
            },{ 
              icon: "fa fa-dollar",
              text: "成果类型",
              id: '015'
            }, 
            { 
              icon: "fa fa-cny",
              text: "科室成果类型",
              id: '016'
            }
          ] 
        },
        {
          text: "权限管理",
          icon: "fa fa-balance-scale",
          // selectedIcon: "glyphicon glyphicon-stop",
          href: "#node-1",
          selectable: true,
          id: '02',
          selectable: false,
          // state: {
            // checked: true,
            // disabled: true,
            // expanded: true,
            // selected: true
          // },
          tags: ['available'],
          nodes: 
          [
            { icon: "fa fa-safari",
              text: '系统权限',
              id: '021'
            },
            { icon: "fa fa-navicon",
              text: '项目权限',
              id: '022'
            }
          ]
        },
        {
          text: "账号管理",
          icon: "fa fa-users icon",
          // selectedIcon: "glyphicon glyphicon-stop",
          href: "#node-1",
          selectable: true,
          id: '03',
          selectable: false,
          // state: {
            // checked: true,
            // disabled: true,
            // expanded: true,
            // selected: true
          // },
          tags: ['available'],
          nodes: 
          [
            { icon: "fa fa-users",
              text: '用户',
              id: '031'
            },
            { icon: "fa fa-th",
              text: 'IP地址段',
              id: '032'
            },
            { icon: "fa fa-group",
              text: '用户组',
              id: '033'
            },
          ]
        }, 
        {
          text: "成果编辑",
          icon: "fa fa-list-alt icon",
          // selectedIcon: "glyphicon glyphicon-stop",
          href: "#node-1",
          selectable: true,
          id: '04',
          selectable: false,
          tags: ['available'],
          nodes: 
          [
            { 
              icon: "fa fa-edit",
              text: "本周成果",
              id: '041'
            },
            { 
              icon: "fa fa-copy",
              text: "本月成果",
              id: '042'
            },
            { 
              icon: "fa fa-lock",
              text: "前月成果",
              id: '043'
            },
            { 
              icon: "fa fa-lock",
              text: "当年成果",
              id: '044'
            }
          ]
        },{
          text: "价值编辑",
          icon: "fa fa-list-alt icon",
          // selectedIcon: "glyphicon glyphicon-stop",
          href: "#node-1",
          selectable: true,
          id: '05',
          selectable: false,
          tags: ['available'],
          nodes: 
          [
            { 
              icon: "fa fa-edit",
              text: "当年价值",
              id: '051'
            },
            { 
              icon: "fa fa-copy",
              text: "去年价值",
              id: '052'
            },
            { 
              icon: "fa fa-lock",
              text: "近5年价值",
              id: '053'
            },
            { 
              icon: "fa fa-lock",
              text: "所有价值",
              id: '054'
            }
          ]
        } 
      ]

      $('#tree').treeview({
        data: data,         // data is not optional
        levels: 2,
        enableLinks: true,
        // multiSelect: true
      });  
        // }
          // alert(JSON.stringify({{.json}}));
         // $('#treeview').treeview('collapseAll', { silent: true });
          // $('#tree').treeview({
          // data: [{{.json}}],//defaultData,
          // data:alternateData,
          // levels: 3,// expanded to 5 levels
          // enableLinks:true,
          // showTags:true,
          // collapseIcon:"glyphicon glyphicon-chevron-up",
          // expandIcon:"glyphicon glyphicon-chevron-down",
        // });

        $('#tree').on('nodeSelected', function(event, data) {
          document.getElementById("iframepage").src="/admin/"+data.id;
          if (data.id=="010"){
            $(".breadcrumb").html("<i class='fa fa-home'></i>后台&gt;日历")
          }else if (data.id=="011"){
            $(".breadcrumb").html("<i class='fa fa-home'></i>后台&gt;系统设置&gt;"+data.text)
          }else if(data.id=="012"){
            $(".breadcrumb").html("<i class='fa fa-home'></i>后台&gt;系统设置&gt;"+data.text)
          }else if(data.id=="013"){
            $(".breadcrumb").html("<i class='fa fa-home'></i>后台&gt;系统设置&gt;"+data.text)
          }else if(data.id=="014"){
            $(".breadcrumb").html("<i class='fa fa-home'></i>后台&gt;系统设置&gt;"+data.text)
          }else if(data.id=="015"){
            $(".breadcrumb").html("<i class='fa fa-home'></i>后台&gt;系统设置&gt;"+data.text)
          }else if(data.id=="016"){
            $(".breadcrumb").html("<i class='fa fa-home'></i>后台&gt;系统设置&gt;"+data.text)
          }else if(data.id=="021"){
            $(".breadcrumb").html("<i class='fa fa-home'></i>后台&gt;权限管理&gt;"+data.text)
          }else if(data.id=="022"){
            $(".breadcrumb").html("<i class='fa fa-home'></i>后台&gt;权限管理&gt;"+data.text)
          }else if(data.id=="031"){
            $(".breadcrumb").html("<i class='fa fa-home'></i>后台&gt;账号管理&gt;"+data.text)
          }else if(data.id=="032"){
            $(".breadcrumb").html("<i class='fa fa-home'></i>后台&gt;账号管理&gt;"+data.text)
          }else if(data.id=="033"){
            $(".breadcrumb").html("<i class='fa fa-home'></i>后台&gt;账号管理&gt;"+data.text)
          }else if(data.id=="041"){
            $(".breadcrumb").html("<i class='fa fa-home'></i>后台&gt;成果编辑&gt;"+data.text)
          }else if(data.id=="042"){
            $(".breadcrumb").html("<i class='fa fa-home'></i>后台&gt;成果编辑&gt;"+data.text)
          }else if(data.id=="043"){
            $(".breadcrumb").html("<i class='fa fa-home'></i>后台&gt;成果编辑&gt;"+data.text)
          }else if(data.id=="051"){
            $(".breadcrumb").html("<i class='fa fa-home'></i>后台&gt;价值编辑&gt;"+data.text)
          }else if(data.id=="052"){
            $(".breadcrumb").html("<i class='fa fa-home'></i>后台&gt;价值编辑&gt;"+data.text)
          }else if(data.id=="053"){
            $(".breadcrumb").html("<i class='fa fa-home'></i>后台&gt;价值编辑&gt;"+data.text)
          }
        }); 

      var obj = {};
      obj.text = "123";
      $("#btn").click(function (e) {
        var arr = $('#tree').treeview('getSelected');
        for (var key in arr) {
          c.innerHTML = c.innerHTML + "," + arr[key].id;
        }
      })
    }) 

    function index1(value,row,index){
    // alert( "Data Loaded: " + index );
      return index+1
    }
</script>

  <div class="col-lg-10">
    <iframe src="/admin/01" name='main' frameborder="0"  width="100%" scrolling="no" marginheight="0" marginwidth="0" id="iframepage" onload="this.height=100"></iframe> 
  </div>  


  <script type="text/javascript">
    function reinitIframe(){//http://caibaojian.com/frame-adjust-content-height.html
      var iframe = document.getElementById("iframepage");
      try{
        var bHeight = iframe.contentWindow.document.body.scrollHeight;
        var dHeight = iframe.contentWindow.document.documentElement.scrollHeight;
        var height = Math.max(bHeight, dHeight); iframe.height = height;
        // console.log(height);//这个显示老是在变化
      }catch (ex){
      } 
    } 
    window.setInterval("reinitIframe()", 200);
  </script>

</body>
</html>