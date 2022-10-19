<?php

namespace Hightemp\WappSnanimageServer\Helpers\HTML;

class BaseHTMLHelper
{
    const T_A = 'a';
    const T_SPAN = 'span';
    const T_DIV = 'div';
    const T_NAV = 'nav';
    const T_UL = 'ul';
    const T_LI = 'li';
    const T_INPUT = 'input';
    const T_TEXTAREA = 'textarea';
    const T_BUTTON = 'button';
    const T_SELECT = 'select';
    const T_FORM = 'from';
    const T_SCRIPT = 'script';
    const T_LINK = 'link';

    public static $aCache = [];

    public static $aDefaultAttrs = [];
    public static $aAttrs = [];

    /** @var bool $bBufferOutput влиет на работу метода fnPrint */
    public static $bBufferOutput = false;
    public static $aBuffer = [];
    public static $aNames = [];
    public static $sCurrentName = "";

    function fnSetValue($mValue)
    {
        $this->aAttrs["value"] = $mValue;
    }

    function fnGetValue($mValue)
    {
        return $this->aAttrs["value"];
    }

    public static function fnBuild()
    {
        if (!isset(static::$aCache[static::class])) {
            static::$aCache[static::class] = new static();
        }
        return static::$aCache[static::class];
    }

    public static function fnCleanBuffer()
    {
        static::$aBuffer = [];
    }

    public static function fnAddToBuffer($mValue)
    {
        static::$aBuffer[static::$sCurrentName][] = $mValue;
    }

    public static function fnBeginBuffer()
    {
        static::$bBufferOutput = true;
        $sName = static::$sCurrentName = microtime();
        static::$aBuffer[$sName] = [];
        static::$aNames[] = $sName;
    }

    public static function fnEndBuffer($bJoin=false)
    {
        $sName = static::$sCurrentName;

        $aOutput = static::$aBuffer[$sName];
        unset(static::$aBuffer[$sName]);

        array_pop(static::$aNames);
        if (static::$aNames) {
            static::$sCurrentName = static::$aNames[count(static::$aNames)-1];
        } else {
            static::$sCurrentName = "";
        }

        if (!count(static::$aBuffer)) {
            static::$bBufferOutput = false;
        }

        return $bJoin ? join("", $aOutput) : $aOutput;
    }

    public static function &fnPrepareAttrs(&$aAttr, $aDefault=[])
    {
        $aAttr = array_replace_recursive($aDefault, $aAttr);
        return $aAttr;
    }

    public static function fnPrepareAttrString($aAttr, $aDefault=[])
    {
        $sResult = "";
        $aAttr = (array) $aAttr;

        if ($aDefault) {
            static::fnPrepareAttrs($aAttr, $aDefault);
        }

        foreach ($aAttr as $sK => $sV) {
            $sV = addslashes($sV);
            $sResult .= "$sK=\"$sV\" \n";
        }

        return "\n".$sResult;
    }

    public static function fnRenderTag($sTagName, $bSingle=false, $aAttrs=[], $sContent="")
    {
        $sAttr = static::fnPrepareAttrString($aAttrs);
        $sHTML = "<".$sTagName." ".$sAttr;
        if ($bSingle) {
            $sHTML .= "/>";
        } else {
            $sHTML .= ">".$sContent."</".$sTagName.">";
        }
        return $sHTML;
    }

    public static function fnPrint($sHTML)
    {
        if (static::$bBufferOutput) {
            static::fnAddToBuffer($sHTML);
        } else {
            echo $sHTML;
        }
    }
}