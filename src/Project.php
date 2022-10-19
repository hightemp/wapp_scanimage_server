<?php 

namespace Hightemp\WappSnanimageServer;

use RedBeanPHP\Facade as R;
use League\Url\Url;

class Project
{
    public static $sCurrentURL = "";
    public static $oCurrentURL = null;

    public static $aHeadCrumbs = [];

    const C_DATABASE = ROOT_PATH.'/data/dbfile.db';
    const C_SCANNED_REL_PATH = "/files/scanned";
    const C_SCANNED_PATH = ROOT_PATH.self::C_SCANNED_REL_PATH;
    const C_SCANNED_FILE_MASK = "*.jpeg";
    const C_ARCHIVED_REL_PATH = "/files/archives";
    const C_ARCHIVED_PATH = ROOT_PATH.self::C_ARCHIVED_REL_PATH;
    const C_ARCHIVED_FILE_MASK = "*.zip";
    const C_CACHE_PATH = ROOT_PATH."/cache";



    public static function fnInit()
    {
        R::setup( 'sqlite:'.static::C_DATABASE );
        
        define("DATE_FORMAT", "Y-m-d H:i:s");
        
        define('T_TABLES', 'tables');
        define('T_FORMS', 'forms');
        define('T_REQUESTS', 'requests');
        define('T_VIEWS', 'views');
        
        static::$sCurrentURL = (isset($_SERVER['HTTPS']) && $_SERVER['HTTPS'] === 'on' ? "https" : "http");
        static::$sCurrentURL .= "://".$_SERVER['HTTP_HOST'].$_SERVER['REQUEST_URI'];
        static::$oCurrentURL = Url::createFromUrl(static::$sCurrentURL);
    }

    public static function fnExecControllerAction()
    {
        $sModule = isset($_GET['m']) && $_GET['m'] ? $_GET['m'] : "Main";
        $sAction = isset($_GET['a']) && $_GET['a'] ? $_GET['a'] : "fnIndexHTML";
        $sControllerClass = "Hightemp\\WappSnanimageServer\\Modules\\{$sModule}\\Controller";

        ob_start();
        $sControllerClass::$sAction();
        $sContent = ob_get_clean(); 

        return $sContent;
    }

    public static function fnLinkTo($sM, $sA, $aAttr=[], $bFullUpdate=true)
    {
        $oCurrentURL = Url::createFromUrl(static::$sCurrentURL);
        if ($bFullUpdate) $oCurrentURL->setQuery('');
        $oQ = $oCurrentURL->getQuery();
        $oQ->modify([
            ...$aAttr,
            'm' => $sM,
            'a' => $sA,
            'redirect' => static::$sCurrentURL
        ]);
        return $oCurrentURL.'';
    }

    public static function fnGetScannedFiles()
    {
        $aFiles = glob(static::C_SCANNED_PATH."/".static::C_SCANNED_FILE_MASK);

        $aFiles = array_map(function($mI) { 
            $aStat = stat($mI);
            return [
                basename($mI),
                $mI,
                static::human_filesize($aStat['size']),
                ...((array)$aStat)
            ]; 
        }, $aFiles);

        return $aFiles;
    }

    public static function fnGetArchivedFiles()
    {
        $aFiles = glob(static::C_ARCHIVED_PATH."/".static::C_ARCHIVED_FILE_MASK);

        $aFiles = array_map(function($mI) { 
            $aStat = stat($mI);
            return [
                basename($mI),
                static::C_ARCHIVED_REL_PATH.'/'.basename($mI),
                static::human_filesize($aStat['size']),
                ...((array)$aStat)
            ]; 
        }, $aFiles);

        return $aFiles;
    }

    public static function fnGetScannersList()
    {
        exec("scanimage -L", $aOutput, $iCode);
        return join("<br>\n", $aOutput);
    }

    public static function fnScanImage()
    {
        $sImageFile = time().".jpeg";
        $sPath = static::fnGetImagePath($sImageFile);
        exec("scanimage --format=jpeg --output-file {$sPath}", $aOutput, $iCode);
        return $sImageFile;
    }

    public static function fnGetScannersListCached()
    {
        $sPath = static::C_CACHE_PATH."/scanners_list.txt";

        if (is_file($sPath)) {
            $sHTML = file_get_contents($sPath);
        } else {
            $sHTML = static::fnGetScannersList();
            file_put_contents($sPath, $sHTML);
        }
        
        return $sHTML;
    }

    public static function fnGetImagePath($sImageFile)
    {
        return static::C_SCANNED_PATH."/".$sImageFile;
    }

    public static function fnGetImageRelPath($sImageFile)
    {
        return static::C_SCANNED_REL_PATH."/".$sImageFile;
    }

    public static function fnGetArchivePath($sArchiveFile)
    {
        return static::C_ARCHIVED_PATH."/".$sArchiveFile;
    }

    public static function fnGetArchiveRelPath($sArchiveFile)
    {
        return static::C_ARCHIVED_REL_PATH."/".$sArchiveFile;
    }

    public static function human_filesize($bytes, $dec = 2) 
    {
        $size   = array('B', 'kB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB');
        $factor = floor((strlen($bytes) - 1) / 3);

        return sprintf("%.{$dec}f", $bytes / pow(1024, $factor)) . @$size[$factor];
    }
}