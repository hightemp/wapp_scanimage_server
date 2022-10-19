<?php

namespace Hightemp\WappSnanimageServer\Helpers\HTML;

use Hightemp\WappSnanimageServer\Helpers\HTML\AlertMessage;

class AlertMessages extends BaseHTMLHelper
{
    public function __invoke($aList)
    {
        $oAlertMessage = AlertMessage::fnBuild();

        foreach ($aList as $aItem) {
            $oAlertMessage($aItem[1], $aItem[0]);
        }
    }
}

