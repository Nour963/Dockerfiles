<!doctype html>
<html>
<head>
<meta charset="utf-8">
<title>WORKING v2!</title>
<style>
body {background-color: #3C69E2;}
h1 {
text-align: center;
margin-top: 250px;
color: white;
}
h3 {
text-align: center;
margin-top: 50px;
color: orange;
font-size: 170%;
}
.container{
position: absolute;
right: 500px;
top: 430px;
width: 13.8%;
}
h2 {
text-align: right;
font-size: 230%;
color: #FEB53F;
margin-top: 0px;
border-style: dotted;
margin-right: 35px;
animation: blinker 1s step-start infinite;
}
@keyframes blinker {
  50% {
    opacity: 0;
  }
}
 .center {
  display: block;
  margin-left: 380px;
  margin-right: auto;
  width: 35%;
  height: 300px;
}
</style>
  
</head>
<body>
<div class='container'>
<h2> version 2 </h2> </div>
  <h1>Congratulations! it's working fine :) you reached the go app.</h1> 

 <IMG class="center" SRC="https://media.giphy.com/media/f5di0y0RLjPe2tHKZF/giphy.gif">

<h3>
<?php
$servername = "localhost";
$username = "root";
$password = "root";

try {
  $conn = new PDO("mysql:host=$servername;dbname=test", $username, $password);
  $conn->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);
  echo "<p  style=\"color:#5aed21;\"> And connected successfully to the database! <p>";
    }

catch(PDOException $e)
    {
    echo "<p  style=\"color:#d80c0c;\"> Could not connect to the database! <br> <br> error: " . $e->getMessage() ."</p>";
    }
?> 
</h3>
</body>
</html>
