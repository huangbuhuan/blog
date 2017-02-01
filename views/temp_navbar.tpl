{{ define "navbar" }}
	<nav class="navbar navbar-default navbar-fixed-top">
	    <div class="container">
	        <a class="navbar-brand" href="/">黄步欢的博客</a>
	        <ul class="nav navbar-nav">
	            <li {{ if .IsHome }}class="active"{{ end }}><a href="/">首页</a></li>
	            <li><a href="">近期文章</a></li>
	            <li><a href="">友情链接</a></li>
	            <li><a href="">计划</a></li>
	            <li><a href="">关于</a></li>
	        </ul>
	    </div>
	</nav>
{{ end }}