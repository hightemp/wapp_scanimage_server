<?php

namespace Hightemp\WappSnanimageServer\Helpers\HTML;

class Input extends BaseHTMLHelper
{
    public static $aDefaultAttrs = [
        "class" => "form-control",
    ];

    public function __invoke($sName, $sValue, $aAttrs=[])
    {
        $aAttrs['name'] = $sName;
        $aAttrs['value'] = $sValue;
        $aAttrs = static::fnPrepareAttrs($aAttrs, static::$aDefaultAttrs);
        static::fnPrint(static::fnRenderTag(static::T_INPUT, true, $aAttrs));
    }
}