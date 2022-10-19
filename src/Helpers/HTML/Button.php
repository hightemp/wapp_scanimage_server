<?php

namespace Hightemp\WappSnanimageServer\Helpers\HTML;

class Button extends BaseHTMLHelper
{
    public static $aDefaultAttrs = [
        "class" => "btn",
        "type" => "submit",
    ];

    public function __invoke($sName, $sTitle, $sValue, $sClassType="secondary", $aAttrs=[])
    {
        $aAttrs['name'] = $sName;
        $aAttrs['value'] = $sValue;
        $aAttrs = static::fnPrepareAttrs($aAttrs, static::$aDefaultAttrs);
        $aAttrs['class'] .= " btn-{$sClassType}";

        static::fnPrint(static::fnRenderTag(static::T_BUTTON, false, $aAttrs, $sTitle));
    }
}

