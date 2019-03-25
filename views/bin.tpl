<!DOCTYPE html>
<html>
<head>
    <title>RequestBingo - {{.bin.Name}}</title>
    <link rel="shortcut icon" href="{{.bin.FaviconUrl}}"/>
    <link href="/static/css/bootstrap.css" rel="stylesheet" media="screen">
    <link href="/static/css/responsive.css" rel="stylesheet" media="screen">
    <link href="/static/css/styles.css" rel="stylesheet" media="screen">
    <link href="/static/css/font-awesome.min.css" rel="stylesheet" media="screen">
    <link type="text/css" href="/static/css/prettify.css" rel="stylesheet"/>
    <script type="text/javascript" src="/static/js/jquery.min.js"></script>
    <script type="text/javascript" src="/static/js/prettify.js"></script>
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
                <a href="/"><span class="logo-accent">Request</span>Bin</a>
            </h1>
            <nav>
                <ul class="nav-menu">
                    <li><a href="/{{.bin.Name}}?inspect"><i class="icon-circle icon-2x"
                                                            style="color: rgb{{.color}}"></i></a>
                        <input type="text" value="{{.base_url}}/{{.bin.Name}}" onclick="this.select()"/>
                        {{if .bin.Private }}<i class="icon-lock"></i>{{end}}
                    </li>
                </ul>
            </nav>
        </div>
    </header>
</div>

<div id="content" class="row-fluid">
    {{$width := 12}}
    {{if .recent}} {{$width = 10}} {{end}}
    <div class="span{{$width}} content-wrap">
        {{range $index, $elem := .bin.Requests}}
        <div class="message-wrapper" id="message-wrapper-{{$elem.Id}}">
            <div class="message-list">
                <div class="row-fluid">
                    <div class="span4">
                        {{$.base_url}}<br>
                        <span class="method">{{$elem.Method}}</span>
                        <span class="absolute-path">{{$elem.Path}}</span><span
                                class="querystring">{{$elem.QueryString}}</span>
                    </div>
                    <div class="span6 content">
                        {{if $elem.ContentType}}<i class="icon-code"></i>{{end}} {{$elem.ContentType}}<br>
                        <i class="icon-cloud-upload"></i> {{$elem.ContentLength}}
                    </div>
                    <div class="span2" class="timestamp">
              <span title="{{$elem.Time}}">{{$elem.Time}} ago
                <a href="#{{$elem.Id}}"><i class="icon-link"></i></a>
              </span><br>
                        From {{$elem.RemoteAddr}}
                    </div>
                </div>
            </div>

            <div id="detail-{{$elem.Id}}" class="message-detail">
                <div id="request-detail-{{$elem.Id}}" class="request-detail">
                    <div class="row-fluid">
                        <div class="span4">
                            <h5>FORM/POST PARAMETERS</h5>
                            {{range $k, $v := $elem.FormData}}
                                <p class="keypair"><strong>{{$k}}:</strong> {{$v}}</p>
                            {{else}}
                                <em>None</em>
                            {{end}}

                            {{if $elem.QueryString}}
                                <h5>QUERYSTRING</h5>
                            {{range $k, $v := $elem.QueryString}}
                            {{if not $v}}
                            <p class="keypair"><strong>{{$k}}</strong></p>
                            {{else}}
                            <p class="keypair"><strong>{{$k}}:</strong> {{$v}}</p>
                            {{end}}
                            {{end}}
                            {{end}}
                        </div>
                        <div class="span8">
                            {{if $elem.Headers}}
                            <h5>HEADERS</h5>
                            {{range $k, $v := $elem.Headers}}
                            <p class="keypair"><strong>{{$k}}:</strong> {{$v}}</p>
                            {{end}}
                            {{end}}
                        </div>
                    </div>

                    <h5>RAW BODY</h5>
                    <div class="request-body" data-id="{{$elem.Id}}">
                        <pre class="body prettyprint">{{if $elem.Raw}}{{$elem.Raw}}{{else}}<em>None</em>{{end}}</pre>
                    </div>
                </div>
            </div>
        </div>
        {{else}}

        <h4 class="text-center">Bin URL</h4>
        <h2 class="text-center">
            <input class="xxlarge input-xxlarge" type="text" value="{{.base_url}}/{{.bin.Name}}" onclick="this.select()"
                   style="border-color: rgb{{.color}}; border-width: 3px;"/></h2>
        <p class="text-center">{{if .bin.Private}}This is a private bin. Requests are only viewable from this
            computer.{{end}}

    <hr>
    <div class="row-fluid">
        <div class="span6 offset3">

            <h4>Make a request to get started.</h4>

            <h5>cURL</h5>
            <pre>curl -X POST -d "fizz=buzz" {{.base_url}}/{{.bin.Name}}</pre>

            <h5>Python (with Requests)</h5>
            <pre class="prettyprint">import requests, time
r = requests.post('{{.base_url}}/{{.bin.Name}}', data={"ts":time.time()})
print r.status_code
print r.content</pre>

            <h5>Node.js (with request)</h5>
            <pre class="prettyprint">var request = require('request');
var url ='{{.base_url}}/{{.bin.Name}}'
request(url, function (error, response, body) {
  if (!error) {
    console.log(body);
  }
});</pre>

            <h5>Ruby</h5>
            <pre class="prettyprint">require 'open-uri'
result = open('{{.base_url}}/{{.bin.Name}}')
result.lines { |f| f.each_line {|line| p line} }</pre>

            <h5>C# / .NET</h5>
            <pre class="prettyprint">using System;
using System.Net.Http;
using System.Threading.Tasks;

namespace RequestBinExample
{
  class Program
  {
    static void Main(string[] args)
    {
      MakeRequest();
    }

    private static async Task MakeRequest()
    {
      var httpClient = new HttpClient();
      var response = await httpClient.GetAsync(new Uri("{{.base_url}}/{{.bin.Name}}"));
      var body = await response.Content.ReadAsStringAsync();
      Console.WriteLine(body);
    }
  }
}</pre>

            <h5>Java</h5>
            <pre class="prettyprint">import org.apache.commons.httpclient.*;
import org.apache.commons.httpclient.methods.*;
import org.apache.commons.httpclient.params.HttpMethodParams;

import java.io.*;

public class RequestBinTutorial {
  public static void main(String[] args) {
    HttpClient client = new HttpClient();
    GetMethod method = new GetMethod("{{.base_url}}/{{.bin.Name}}");
    try {
      int statusCode = client.executeMethod(method);
      byte[] responseBody = method.getResponseBody();
      System.out.println(new String(responseBody));
    } catch (Exception e) {
      System.err.println("Fatal error: " + e.getMessage());
      e.printStackTrace();
    } finally {
      method.releaseConnection();
    }
  }
}</pre>

            <h5>PHP</h5>
            <pre class="prettyprint">&lt;?php
    $result = file_get_contents('{{.base_url}}/{{.bin.Name}}');
    echo $result;
?&gt;</pre>

        </div>
    </div>

        {{end}}}

        <hr>

        <div class="alert-message block-message info">
            <h4>Limits</h4>
            <p>This {{if .bin.Private }}<strong>private</strong>{{end}}
                bin will keep the last 20 requests made to it and remain available for 48 hours after it was created.
                However, data might be cleared at any time, so <strong>treat bins as highly ephemeral</strong>.</p>

            <h4>Need more?</h4>
            <p><a href="https://www.runscope.com/?utm_source=partner&amp;utm_medium=link&amp;utm_campaign=requestbin">Sign
                    up for a free Runscope account</a>. Runscope request captures give you live updates, permanent URLs,
                file handling and much more.</p>
        </div>
    </div>
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
