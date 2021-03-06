<!DOCTYPE html>
<html>
<head>
    <title>RequestBingo &mdash; Collect, inspect and debug HTTP requests and webhooks</title>
    <link href="/static/img/logo.png" rel="shortcut icon"/>
    <link href="/static/css/bootstrap.css" rel="stylesheet" media="screen">
    <link href="/static/css/responsive.css" rel="stylesheet" media="screen">
    <link href="/static/css/styles.css" rel="stylesheet" media="screen">
    <link href="/static/css/font-awesome.min.css" rel="stylesheet" media="screen">
    <link type="text/css" href="/static/css/prettify.css" rel="stylesheet"/>
    <script type="text/javascript" src="/static/js/jquery.min.js"></script>
    <script type="text/javascript" src="/static/js/prettify.js"></script>
    <script type="text/javascript">
        function createBin() {
            $.ajax({
                'url': '/api/v1/bins', 'type': 'POST',
                'data': {'private': $('#private').prop("checked")},
                'success': function (data) {
                    window.location.href = "/" + data['name'] + "?inspect";
                }
            });
        }
    </script>
</head>
<body onload="prettyPrint()">

<div class="runscope-wrap">
    <div class="row-fluid">
        <p class="tagline">A <a
                    href="https://www.runscope.com?utm_source=partner&amp;utm_medium=link&amp;utm_campaign=requestbin"
                    title="Runscope API Tools"><strong>Runscope</strong></a> community project. <a
                    href="mailto:requestbin@runscope.com">Send us feedback!</a></p>
    </div>
</div>
<div class="header-wrap">
    <header class="row-fluid">
        <div class="span12">
            <h1 class="logo">
                <a href="/"><img src="/static/img/logo-2x.png" height="38" width="44"/></a>
                <a href="/"><span class="logo-accent">Request</span>Bingo</a>
            </h1>
            <nav>
                <ul class="nav-menu">
                    <li></li>
                </ul>
            </nav>
        </div>
    </header>
</div>

<div id="content" class="row-fluid">
    {{$width := 12}}
    {{if .recent}} {{$width = 10}} {{end}}
    <div class="span{{$width}} content-wrap">
        <div class="row-fluid">
            <div class="banner-hero">
                <h2 class="banner-header"><span class="accent-bg-red">Inspect HTTP Requests</span></h2>
                <h3 class="banner-subheader">
                    RequestBin gives you a URL that will collect requests made to it and let you inspect them in a
                    human-friendly way.<br>
                    Use RequestBin to see what your HTTP client is sending or to inspect and debug webhook requests.
                </h3>

                <form class="form-inline">
                    <p class="banner-button">
                        <a class="btn btn-success btn-large" onclick="createBin()"><i class="icon-plus-sign"></i> Create
                            a RequestBin</a>
                    </p>

                    <p class="banner-option">
                        <label class="checkbox" title="Private bins can only be viewed by you using a browser cookie">
                            <input id="private" type="checkbox"/> &nbsp;Private
                            <small>(only viewable from this browser)</small>
                        </label>
                    </p>
                </form>
            </div>
        </div>
    </div>

    {{if .recent}}
    <div class="span2">
        <h5><i class="icon-time"></i> History</h5>
        <ul id="recent">
            {{range $index, $elem := .recent}}
            <li>
                <a href="/{{$elem.Name}}?inspect"><i class="icon-circle" style="color: rgb{{$elem.ColorString}}"></i></a>
                <a href="/{{$elem.Name}}?inspect">{{$elem.Name}}</a>
                {{if $elem.Private}}<i class="icon-lock"></i>{{end}}
                ({{$elem.Requests | len}})
            </li>
            {{else}}}
            <p style="margin-left: -24px; color: gray;">
                No recent bins.
            </p>
            {{end}}
        </ul>
    </div>
    {{end}}
</div>

<div class="footer-wrap">
    <footer class="row-fluid">
        <div class="span12">
            <div class="footer-logo">
                <img src="/static/img/logo-runscope-2x.png" height="34" width="34">
            </div>
            <h3 class="footer-slogan">
                RequestBin is brought to you by Runscope.<br>Solve API errors fast. Debug, test and share your API
                calls.
                <br>
                <a href="https://www.runscope.com/?utm_source=partner&amp;utm_medium=link&amp;utm_campaign=requestbin">Learn
                    more <i class="icon-long-arrow-right"></i></a>
            </h3>

            <p class="footer-hero">
                <a href="https://www.runscope.com/?utm_source=partner&amp;utm_medium=link&amp;utm_campaign=requestbin"
                   title="Learn More"><img src="/static/img/runscope-hero.png"></a>
            </p>

            <ul class="footer-menu">
                <li class="footer-link footer-title">Runscope</li>
                <li class="footer-link"><a
                            href="https://www.runscope.com/?utm_source=partner&amp;utm_medium=link&amp;utm_campaign=requestbin"
                            title="Learn More">Learn More</a></li>
                <li class="footer-link"><a href="http://blog.runscope.com" title="Runscope Blog">Blog</a></li>
                <li class="footer-link"><a href="https://twitter.com/Runscope" title="Follow Runscope on Twitter">Twitter</a>
                </li>
                <li class="footer-link"><a href="https://plus.google.com/110044295418761600228"
                                           rel="publisher">Google+</a></li>
                <li class="footer-link"><a href='https://alpha.app.net/runscope' rel='me'>App.net</a></li>
            </ul>
            <p class="copyright">RequestBin originally created by <a href="http://progrium.com">Jeff Lindsay</a>.</p>
            <p class="copyright">&copy; 2013 Runscope Inc. - <a
                        href="https://www.runscope.com/privacy?utm_source=partner&amp;utm_medium=link&amp;utm_campaign=requestbin">Privacy
                    Policy</a> - <a
                        href="https://www.runscope.com/terms?utm_source=partner&amp;utm_medium=link&amp;utm_campaign=requestbin">Terms
                    of Service</a></p>
        </div>
    </footer>
</div>
</body>
</html>
