<?php

namespace Hightemp\WappSnanimageServer\Helpers\HTML;

use Hightemp\WappSnanimageServer\Helpers\HTML\BaseHTMLHelper;

class A extends BaseHTMLHelper
{
    public function __invoke($sContent, $sHref, $aAttrs=[])
    {
        $aAttrs = static::fnPrepareAttrs($aAttrs);
        $aAttrs['href'] = $sHref;
        static::fnPrint(static::fnRenderTag(static::T_A, false, $aAttrs, $sContent));
    }
}