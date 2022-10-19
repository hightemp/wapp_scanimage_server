<?php

namespace Hightemp\WappSnanimageServer;

include_once(ROOT_PATH."/RedBeanPHP.php");

Project::fnInit();

$sContent = Project::fnExecControllerAction();

require_once("layout.php");