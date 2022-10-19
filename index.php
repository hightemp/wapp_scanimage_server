<?php

define('ROOT_PATH', __DIR__);
define('DEBUG', getenv('DEBUG'));
define('SERVER_PORT', getenv('SERVER_PORT') ?: 8080);
define('SERVER_HOST', getenv('SERVER_HOST') ?: "0.0.0.0");

if (defined('DEBUG') && DEBUG) {
    ini_set('display_errors', 1);
    error_reporting(E_ALL & ~E_DEPRECATED);
}

require_once("vendor/autoload.php");
require_once("src/index.php");