<?php

namespace Hightemp\WappSnanimageServer\Helpers\HTML;

class AlertMessage extends BaseHTMLHelper
{
    public static $aDefaultAttrs = [
        "class" => "alert",
        "role" => "alert",
    ];

    public function __invoke($sMessage, $sClassType="secondary", $aAttrs=[])
    {
        if ($sClassType) {
            static::$aDefaultAttrs['class'] .= " alert-{$sClassType}";
        }

        $aAttrs = static::fnPrepareAttrs($aAttrs, static::$aDefaultAttrs);
        static::fnPrint(static::fnRenderTag(static::T_DIV, false, $aAttrs, $sMessage));
    }
}

