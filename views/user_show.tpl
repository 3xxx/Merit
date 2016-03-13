<!-- 测试页面 -->
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


<div id="treeview" class="col-xs-3"></div>

<div class="col-lg-9">
  <table class="table table-striped">
    <thead>
      <tr>
        <th>#序号</th>
        <th>单位</th>
        <th>分院</th>
        <th>科室</th>
        <th>价值分类</th>
        <th>价值项</th>
        <th>选择项</th>
        <th>操作</th>
      </tr>
    </thead>

    <tbody>
      <tr>
        <th></th>
        <th>{{.Input.Danwei}}</th>
        <th></th>
        <th></th>
        <th></th>
        <th></th>
        <th></th> 
        <th></th>
        {{range $k0,$v0 :=$.Input.Fenyuan}}
                  <tr>
                  <th></th>
                  <th></th>
                  <th>{{.Department}}</th>
                  <th></th>
                  <th></th>
                  <th></th> 
                  <th></th>
                  <th></th>
                  </tr>
            {{range $k0,$v0 :=.Bumen}}
                  <tr>
                  <th></th>
                  <th></th>
                  <th></th>
                  <th>{{.Keshi}}</th>
                  <th></th>
                  <th></th> 
                  <th></th>
                  <th></th>
                  </tr>

                {{range $k0,$v0 :=.Kaohe}}
                    <tr>
                    <th></th>
                    <th></th>
                    <th></th>
                    <th></th>
                    <th>{{.Category}}</th>
                    <th></th> 
                    <th></th>
                    <th></th>
                    </tr>

                    {{range $k0,$v0 :=.Fenlei}}
                      <tr>
                      <th></th>
                      <th></th>
                      <th></th>
                      <th></th>
                      <th></th> 
                      <th>{{.Project}}</th>
                      <th></th>
                      <th>
                        <a href="/">显示成果</a>
                        <a href="/">修改</a>
                        <a href="/">删除</a></th>
                      </tr>

                      {{range $k0,$v0 :=.Xuanze}}
                        <tr>
                        <th></th>
                        <th></th>
                        <th></th>
                        <th></th>
                        <th></th>
                        <th></th>
                        <th>{{.Choose}}</th>
                        <th></th>
                        </tr>

                      {{end}}

                    {{end}}

                {{end}}
            {{end}}    
        {{end}}
      </tr>

    </tbody>
  </table>
</div>

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
            "text": 'Parent 1',
            "selectable":false,
            // icon: "glyphicon glyphicon-stop",
            // selectedIcon: "glyphicon glyphicon-heart",
            href: '#parent1',
            tags: ['4'],
            // state: {
            // checked: true,
            // disabled: false,
            // expanded: false,
            // selected: true
            // },
            // tags: ['available'],
            nodes: [
              {
                text: 'Child 1',
                selectable:false,
                // icon: "glyphicon glyphicon-stop",
                // selectedIcon: "glyphicon glyphicon-heart",                
                // href: '#child1',
                tags: [2,3],
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
             state: {
            checked: true,
            disabled: false,
            expanded: false,
            selected: true
            },           
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

        var json=[
        {
  "text": "省水利设计分院",
  "selectable": false,
  "nodes": [
    {
      "text": "施工预算分院",
      "selectable": false,
             state: {
            checked: false,
            disabled: false,
            expanded: false,
            selected: false
            },      
      "nodes": [
        {
          "text": "水工室",
          "selectable": false,
          "nodes": [
            {
              "text": "项目管理类",
              "selectable": false,
              "tags": [
                4,
                2
              ],
              "nodes": [
                {
                  "text": "项目负责人",
                  "href": "/add?id=165",
                  "tags": [
                    3,
                    1
                  ],
                  "Mark2": "",
                  "Xuanze": ",大型,中型,小型",
                  "Mark1": ",4,3,2"
                },
                {
                  "text": "课题研究",
                  "href": "/add?id=166",
                  "tags": [
                    3,
                    1
                  ],
                  "Mark2": "4",
                  "Xuanze": "",
                  "Mark1": ""
                }
              ],
              "Parent2": ""
            },
            {
              "text": "贡献类",
              "tags": [
                4,
                2
              ],
              "nodes": [
                {
                  "text": "获奖",
                  "href": "/add?id=168",
                  "tags": [
                    3,
                    1
                  ],
                  "Mark2": "",
                  "Xuanze": ",国家级,省级,院级",
                  "Mark1": ",4,3,2"
                },
                {
                  "text": "开发",
                  "href": "/add?id=169",
                  "tags": [
                    3,
                    1
                  ],
                  "Mark2": "",
                  "Xuanze": ",系统级,标准",
                  "Mark1": ",5,2"
                }
              ],
              "Parent2": ""
            }
          ]
        },
        {
          "text": "施工室",
          "selectable": false,
          "nodes": [
            {
              "text": "项目管理类",
              "tags": [
                4,
                2
              ],
              "nodes": [
                {
                  "text": "施工负责人",
                  "href": "/add?id=172",
                  "tags": [
                    3,
                    1
                  ],
                  "Mark2": "4",
                  "Xuanze": "",
                  "Mark1": ""
                },
                {
                  "text": "课题a",
                  "href": "/add?id=173",
                  "tags": [
                    3,
                    1
                  ],
                  "Mark2": "2",
                  "Xuanze": "",
                  "Mark1": ""
                }
              ],
              "Parent2": ""
            }
          ]
        }
      ]
    },
    {
      "text": "水工分院",
      "selectable": false,
      "nodes": [
        {
          "text": "水工室",
          "selectable": false,
          "nodes": [
            {
              "text": "项目管理类",
              "tags": [
                4,
                2
              ],
              "nodes": [
                {
                  "text": "项目负责人",
                  "href": "/add?id=177",
                  "tags": [
                    3,
                    1
                  ],
                  "Mark2": "",
                  "Xuanze": ",大型,中型,小型",
                  "Mark1": ",6,4,2"
                }
              ],
              "Parent2": ""
            },
            {
              "text": "贡献类",
              "tags": [
                4,
                2
              ],
              "nodes": [
                {
                  "text": "获奖",
                  "href": "/add?id=179",
                  "tags": [
                    3,
                    1
                  ],
                  "Mark2": "",
                  "Xuanze": ",国家级,省级,院级",
                  "Mark1": ",4,3,2"
                },
                {
                  "text": "开发",
                  "href": "/add?id=180",
                  "tags": [
                    3,
                    1
                  ],
                  "Mark2": "",
                  "Xuanze": ",系统级,标准",
                  "Mark1": ",5,2"
                }
              ],
              "Parent2": ""
            }
          ]
        }
      ]
    }
  ]
}
        ];
          // $('#treeview').treeview('collapseAll', { silent: true });
          $('#treeview').treeview({
          data: json,//defaultData,
          // data:alternateData,
          enableLinks:true,
          showTags:true,
          // collapseIcon:"glyphicon glyphicon-chevron-up",
          // expandIcon:"glyphicon glyphicon-chevron-down",
        });
});

</script>

</body>
</html>
