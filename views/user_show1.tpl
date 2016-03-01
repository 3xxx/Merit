<!DOCTYPE html>

<html>
<head>
 <meta charset="UTF-8">
  <title>技术人员价值评测系统</title>
<script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
 <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
 <script src="/static/js/bootstrap-treeview.js"></script>
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
</head>


<div id="treeview" class="col-xs-2"></div>

<div class="col-lg-10">
  <table class="table table-striped">
    <thead>
      <tr>
        <th>#序号</th>
        <th>部门</th>
        <th>价值分类</th>
        <th>选择项</th>
        <th>~</th>
        <th>~</th>
        <th>操作</th>
      </tr>
    </thead>

    <tbody>
      {{range $k,$v :=.Input}}
      <tr>
        <th>{{$k}}</th>
        {{range $k0,$v0 :=.Father}}
        <th>{{.Department}}</th>
                   <th></th>
                   <th></th>
                   <th></th> 
                   <th></th>
                   <th></th>

                 {{range $k1,$v1 :=$.Input}}
                 {{range $k2,$v2 :=.Name}}
                 {{if eq $v2.Parent2 $v0.Department}}
                 <tr>
                   <th></th>
                   <th></th>
                 <th>{{.Category}}</th>
                 <th></th>
                 <th></th>
                 <th></th>
                 <th></th>
                 </tr>


                 {{range $k3,$v3 :=$.Input}}
                 {{range $k4,$v4 :=.List}}
                 {{if eq $v4.Parent $v2.Category}}
                 {{if eq $v4.Grand $v0.Department}}
                 <tr>
                   <th></th>
                   <th></th>
                   <th></th>
                  <th><a href="/">{{.Choose}}</a></th>
                  <th></th>  
                  <th></th>                
                  <th>
                  <a href="/">添加</a>
                  <a href="/">修改</a>
                  <a href="/">删除</a>
                  </th>
                 </tr>
                 {{end}} 
                 {{end}}
                 {{end}} 
                 {{end}} 
                 {{end}}
                 {{end}} 
                 {{end}}
                 {{end}}
      </tr>
      {{end}}

    </tbody>
  </table>
</div>
<!-- {{range $k1,$v1 :=.Input}}
         <tr><th><a href="/" id="name">{{.Father}}</a></th> </tr>
         {{range .Name}}
         <tr><th><a href="/" id="name">{{.Category}}</a></th></tr>
        {{end}}
        {{range .List}}
         <tr><th><a href="/" id="name">{{.Choose}}</a></th></tr>
        {{end}}     
{{end}} -->
<button type="button" class="btn btn-primary btn-lg" style="color: rgb(212, 106, 64);">
<span class="glyphicon glyphicon-user"></span> User
</button>

<button type="button" class="btn btn-primary btn-lg" style="text-shadow: black 5px 3px 3px;">
<span class="glyphicon glyphicon-user"></span> User
</button>
<script type="text/javascript">
$(function() {
        var defaultData = [
          {
            text: 'Parent 1',
            // icon: "glyphicon glyphicon-stop",
            // selectedIcon: "glyphicon glyphicon-heart",
            href: '#parent1',
            tags: ['4'],
            state: {
            checked: true,
            disabled: false,
            expanded: false,
            selected: true
            },
            tags: ['available'],
            nodes: [
              {
                text: 'Child 1',
                // icon: "glyphicon glyphicon-stop",
                // selectedIcon: "glyphicon glyphicon-heart",                
                href: '#child1',
                tags: ['2'],
                nodes: [
                  {
                    text: 'Grandchild 1',
                    href: '#grandchild1',
                    tags: ['0']
                  },
                  {
                    text: 'Grandchild 2',
                    href: '#grandchild2',
                    tags: ['0']
                  }
                ]
              },
              {
                text: 'Child 2',
                href: '#child2',
                tags: ['0']
              }
            ]
          },
          {
            text: 'Parent 2',
            href: '#parent2',
            tags: ['0'],
            nodes: [
              {
                text: 'Child 1',
                href: '#child1',
                tags: ['2'],
                nodes: [
                  {
                    text: 'Grandchild 1',
                    href: '#grandchild1',
                    tags: ['0']
                  },
                  {
                    text: 'Grandchild 2',
                    href: '#grandchild2',
                    tags: ['0']
                  }
                ]
              },
              {
                text: 'Child 2',
                href: '#child2',
                tags: ['0']
              }
            ]
          },
          {
            text: 'Parent 3',
            href: '#parent3',
             tags: ['0']
          },
          {
            text: 'Parent 4',
            href: '#parent4',
            tags: ['0']
          },
          {
            text: 'Parent 5',
            href: '#parent5'  ,
            tags: ['0']
          }
        ];

        var alternateData = [
          {
            text: 'Parent 1',
            tags: ['2'],
            nodes: [
              {
                text: 'Child 1',
                tags: ['3'],
                nodes: [
                  {
                    text: 'Grandchild 1',
                    tags: ['6']
                  },
                  {
                    text: 'Grandchild 2',
                    tags: ['3']
                  }
                ]
              },
              {
                text: 'Child 2',
                tags: ['3']
              }
            ]
          },
          {
            text: 'Parent 2',
            tags: ['7']
          },
          {
            text: 'Parent 3',
            icon: 'glyphicon glyphicon-earphone',
            href: '#demo',
            tags: ['11']
          },
          {
            text: 'Parent 4',
            icon: 'glyphicon glyphicon-cloud-download',
            href: '/demo.html',
            tags: ['19'],
            selected: true
          },
          {
            text: 'Parent 5',
            icon: 'glyphicon glyphicon-certificate',
            color: 'pink',
            backColor: 'red',
            href: 'http://www.tesco.com',
            tags: ['available','0']
          }
        ];
          // $('#treeview').treeview('collapseAll', { silent: true });
          $('#treeview').treeview({
          data: [{{.json}}],//defaultData,
          // collapseIcon:"glyphicon glyphicon-chevron-up",
          // expandIcon:"glyphicon glyphicon-chevron-down",
        });
});

// function getTree() {
//   // Some logic to retrieve, or generate tree structure
//   var tree = [
//   {
//     text: "Parent 1",
//     nodes: [
//       {
//         text: "Child 1",
//         nodes: [
//           {
//             text: "Grandchild 1"
//           },
//           {
//             text: "Grandchild 2"
//           }
//         ]
//       },
//       {
//         text: "Child 2"
//       }
//     ]
//   },
//   {
//     text: "Parent 2"
//   },
//   {
//     text: "Parent 3"
//   },
//   {
//     text: "Parent 4"
//   },
//   {
//     text: "Parent 5"
//   }
// ];
//   return data;
// }

// $('#tree').treeview({data: getTree()});
</script>


<!-- <body>
<div style="text-align:center;clear:both">

</div>

  <ul id="accordion" class="accordion">
    <li>
      <div class="link"><i class="fa fa-paint-brush"></i>Diseño web<i class="fa fa-chevron-down"></i></div>
      <ul class="submenu">
        <li><a href="#">Photoshop</a></li>
        <li><a href="#">HTML</a></li>
        <li><a href="#">CSS</a></li>
        <li><a href="#">Maquetacion web</a></li>
      </ul>
    </li>
    <li>
      <div class="link"><i class="fa fa-code"></i>Desarrollo front-end<i class="fa fa-chevron-down"></i></div>
      <ul class="submenu">
        <li><a href="#">Javascript</a></li>
        <li><a href="#">jQuery</a></li>
        <li><a href="#">Frameworks javascript</a></li>
      </ul>
    </li>
    <li>
      <div class="link"><i class="fa fa-mobile"></i>Diseño responsive<i class="fa fa-chevron-down"></i></div>
      <ul class="submenu">
        <li><a href="#">Tablets</a></li>
        <li><a href="#">Dispositivos mobiles</a></li>
        <li><a href="#">Medios de escritorio</a></li>
        <li><a href="#">Otros dispositivos</a></li>
      </ul>
    </li>
    <li><div class="link"><i class="fa fa-globe"></i>Posicionamiento web<i class="fa fa-chevron-down"></i></div>
      <ul class="submenu">
        <li><a href="#">Google</a></li>
        <li><a href="#">Bing</a></li>
        <li><a href="#">Yahoo</a></li>
        <li><a href="#">Otros buscadores</a></li>
      </ul>
    </li>
  </ul>
  <script src="/static/js/cebianlan.js"></script> -->



<!-- <aside class="accordion">
<h1>News</h1>
<div class="opened-for-codepen">
<h2>News Item #1</h2>
<div class="opened-for-codepen">
<h3>News Item #1a</h3>
<div>
<h4>News Subitem 1</h4>
<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,  quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</p>

<h4>News Subitem 2</h4>
<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,  quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</p>

<h4>News Subitem 3</h4>
<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,  quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</p>
</div>

<h3>News Item #1b</h3>
<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,  quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</p>

<h3>News Item #1c</h3>
<div class="opened-for-codepen">
<h4>News Subitem 1</h4>
<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,  quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</p>

<h4>News Subitem 2</h4>
<p class="opened-for-codepen">Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. </p>
</div>
</div>

<h2>News Item #2</h2>
<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</p>

<h2>News Item #3</h2>
<div>
<h3>News Item #3a</h3>
<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,  quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</p>

<h3>News Item #3b</h3>
<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,  quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</p>
</div>
</div>

<h1>Updates</h1>
<div>
<h2>Update #1</h2>
<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,  quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</p>

<h2>Update #2</h2>
<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</p>

<h2>Update #3</h2>
<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</p>

<h2>Update #4</h2>
<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</p>
</div>

<h1>Miscellaneous</h1>
<div>
<h2>Misc. #1</h2>
<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,  quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</p>

<h2>Misc. #2</h2>
<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</p>

<h2>Misc. #3</h2>
<div>
<h3>Misc. Item #1a</h3>
<div>
<h4>Misc. Subitem 1</h4>
<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,  quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</p>

<h4>Misc. Subitem 2</h4>
<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,  quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</p>

<h4>Misc. Subitem 3</h4>
<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,  quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</p>
</div>
<h3>Misc. Item #2a</h3>
<div>
<h4>Misc. Subitem 1</h4>
<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,  quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</p>

<h4>Misc. Subitem 2</h4>
<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,  quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</p>
</div>
</div>
</div>
</aside>

<script src="/static/js/celan.js"></script> -->
</body>

</html>

<!-- ajax操作json实现动态下拉列表 (2011-12-03 01:53:52)转载▼
标签： 杂谈  
先给出一段兼容浏览器的获取AJAX对象的javascript函数
function getXmlHttp(){
var Xmlhttp=null;
try{
Xmlhttp = new ActiveXObject("MSXML2.XMLHTTP");
}catch(e){
try{Xmlhttp = new ActiveXObject("Microsoft.XMLHTTP");
}catch(E){
Xmlhttp = false;}
}
if((!Xmlhttp) && (typeof(XMLHttpRequest)!= 'undefined') ) {
Xmlhttp = new XMLHttpRequest();
}
return Xmlhttp;
}
使用方法 var xmlhttp = getXmlHttp();
在很多的ajax范例中,开发者都是用xmlhttp从服务器端获得一个xml数据,然后转换成javascript可触及的对象,再用js绘制到document中.但我觉得这并非唯一选项,我甚至觉得是多此一举!为什么不直接传递js对象呢?在我开发的系统中,xmlhttp从服务器上获得的是代表js对象的字符串.假如我要传送一个人员列表,我会在服务器上输出:
[{id:1,name:"name1"},{id:2,name:"name2"},{id:3,name:"name3"}]
然后在浏览器上用js获得这个字符串所代表的对象:
var returned = xmlhttp.responseText;
var obj = eval(returned );
接着,你就可以这样访问:
var person1 = obj[0]; var person2 = obj[1];
alert(person1.id);
alert(person1.name);
这样做比传递xml文档直接一些,不必通过转换可以让js直接访问数据.
ajax操作json举例: 使用JSON来做数据传输的动态下拉列表
动态下拉列表的原理其实很简单的,当某一下拉列表触发了onchange事件,然后使用AJAX在后台向服务器异步请求数据,最后将服务器返回的数据进行解析后动态添加到指定的select上即可!
首先来看后台的数据输出,我们假设服务器传送给客户段的JSON数据格式为如下:
{
"options" : [
{"value" : 值,"text" : 文本},
{"value" : 值,"text" : 文本},
{"value" : 值,"text" : 文本}
]
}
其中options是整个JSON对象的标识符,它是一个数组,该数组中的每一个值表示一个select中的option,当然该值也是一个对象了,有两个属性,一个是value,一个是text,分别对应option中的value和显示的text值.
有了数据格式,那么客户端和服务器端进行交流就方便很多了.我们来先写客户端的JS方法.这里我是提供一个静态的实用类Select,专门针对select元素的操作方法,如下:

function Select(){};

Select.create = function( selectId, json ) {
Select.clear(selectId);
Select.add(selectId, json);
};

Select.add = function(selectId,json) {
try {
if (!json.options) return;
for (var i = 0; i < json.options.length; i ++) {
Select.addOption(selectId,json.options[i].value,json.options[i].text);
}
} catch (ex) {
base.alert('设置select错误:指定的JSON对象不符合Select对象的解析要求!');
}
};

Select.createOption = function( value, text) {
var opt = document.createElement_x('option');
opt.setAttribute('value', value);
opt.innerHTML = text;
return opt;
};

Select.addOption = function( selectId, value, text) {
var opt = Select.createOption(value, text);
$(selectId).appendChild(opt);
return opt;
};

Select.getSelected = function( selectId) {
var slt = $(selectId);
if (!slt) return null;
if (slt.type.toLowerCase() == "select-multiple") {
var len = Select.len(selectId);
var result = [];
for (var i = 0; i < len; i ++) {
if (slt.options[i].selected) result.push(slt.options[i]);
}
return result.length > 1 ? result : (result.length == 0 ? null : result[0]);
} else {
var index = $(selectId).selectedIndex;
return $(selectId).options[index];
}
};

Select.select = function( selectId, index) {
var slt = $(selectId);
if (!slt) return false;
for (var i = 0; i < slt.options.length; i ++) {
if (index == i) {
slt.options[i].setAttribute("selected", "selected");
return true;
}
}
return false;
};

Select.selectAll = function( selectId) {
var len = Select.len(selectId);
for (var i = 0; i < len; i ++) Select.select(selectId, i);
};

Select.len = function( selectId) {
return $(selectId).options.length;
};

Select.clear = function( selectId, iterator) {
if (typeof(iterator) != 'function') {
$(selectId).length = 0;
} else {
var slt = $(selectId);
for (var i = slt.options.length - 1; i >= 0; i --) {
if (iterator(slt.options[i]) == true) slt.removeChild(slt.options[i]);
}
}
};

Select.copy = function( srcSlt, targetSlt, iterator) {
var s = $(srcSlt), t = $(targetSlt);
for (var i = 0; i < s.options.length; i ++) {
if (typeof(iterator) == 'function') {
if (iterator(s.options[i], $(targetSlt).options) == true) {
t.appendChild(s.options[i].cloneNode(true));
}
} else {
t.appendChild(s.options[i].cloneNode(true));
}
}
};
那么在回调方法中就可以只要来调用:
……
var jsonString = xmlHttp.responeText; // 获取服务器返回的json字符串
var jsonObj = null;
try {
jsonObj = eval_r('(' + jsonString + ')'); // 将json字符串转换成对象
} catch (ex) {
return null;
}
Select.create("你的select的ID", jsonObj); // 执行option的添加
……
在Select中提供了很多实用的JS方法来方便操作select对象,我们这里只使用其中的create方法.客户端有了JS,我们现在服务器端的字符串输出json数据.
这里我们用到了JSONLib库,该库可以很方便的来从现有的javaBean或其他集合对象中来转换成json字符串.我们这里也提供一个公用类如下:
package common.utils.json;
import java.util.HashMap;
import java.util.Iterator;
import java.util.List;
import java.util.Map;
import net.sf.json.JSONArray;
import net.sf.json.JSONObject;
import org.apache.commons.beanutils.BeanUtils;
import org.apache.log4j.Logger;

public class SelectJSON
{

private static final Logger log = Logger.getLogger(SelectJSON.class);

public static String fromMap(Map map)
{
Iterator it = map.keySet().iterator();
JSONArray options = new JSONArray();
while (it.hasNext()) {
JSONObject option = new JSONObject();
String key = (String)it.next();
option.put("value", key);
option.put("text", map.get(key));
options.put(option);
}
JSONObject result = new JSONObject();
result.put("options", options.toString());
return result.toString();
}


public static String fromList(List list, String valueProp, String textProp)
{
JSONArray options = new JSONArray();
try {
for (Object obj : list) {
JSONObject option = new JSONObject();
String value = BeanUtils.getProperty(obj, valueProp);
String text = BeanUtils.getProperty(obj, textProp);
option.put("value", value);
option.put("text", text);
options.put(option);
}
} catch (Exception ex) {
throw new RuntimeException(ex);
}
JSONObject result = new JSONObject();
result.put("options", options.toString());
return result.toString();
}

public static void main(String[] args)
{
// map 测试
Map<String,String> tt = new HashMap<String,String>();
tt.put("value1", "text1");
tt.put("value2", "text2");
tt.put("value3", "text3");
log.info(SelectJSON.fromMap(tt));
}
}
在类SelectJSON中提供了两个方法,一个是从map中来获取并转换成json字符串,还一个就是从list中的对象来获取,这个方法需要使用BeanUtils工具来辅助获取JavaBean对象中的指定属性.当然了,你可以可以写其他方便发辅助方法来达到这样的效果.
比如我们在上面的SelectJSON类中的测试,会输出:
{"options":[{"value":"value1","text":"text1"},{"value":"value2","text":"text2"},{"value":"value3","text":"text3"}]}
然后我们再调用上面提到的JS类Select进行操作就可以了,注意,在使用Select类进行操作前,比如先转换服务器返回的字符串为js对象,即使用eval来执行返回字符串即可! -->