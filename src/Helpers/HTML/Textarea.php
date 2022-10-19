<?php

namespace Hightemp\WappSnanimageServer\Helpers\HTML;

class Textarea extends BaseHTMLHelper
{
    public static $aDefaultAttrs = [
        "class" => "form-control",
        "rows" => "10",
    ];

    public function __invoke($sName, $sValue, $aAttrs=[])
    {
        $aAttrs['name'] = $sName;
        $aAttrs = static::fnPrepareAttrs($aAttrs, static::$aDefaultAttrs);
        static::fnPrint(static::fnRenderTag(static::T_TEXTAREA, false, $aAttrs, $sValue));
    }
}